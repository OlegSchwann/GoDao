<h3>Описание</h3>

GoDao — легковесный кодогенератор для упрощения работы с сырым PostgreSQL. Он создан что бы предотвратить написание
шаблонного кода раз за разом, для сериализации и распаковки данных. Вы только пишете сигнатуру функции, а тело
реализуете на SQL, со всеми его достоинствами. Клеевойкод будет сгенерирован за вас.

<h3>Установка</h3>

```bash
# install
go get -u github.com/OlegSchwann/GoDao/...

# use
GoDao <input_file>.go
```

<h3>Синтаксис и примеры</h3>

Идея очень проста — пишется структура, с полями функциями. Теги полей являются кодом на SQL. Если вы пользуетесь GoLand,
то оцените решение — по Сtrl+MouseButton на использовании функции вы сразу перейдёте на объявление SQL, без
необходимости искать строки по всему проекту, сваленные кучей, как это бывает иногда. Простой пример, традиционный для
первого знакомства:

```go
package example

// GoDao: generate
type goDao struct {
   // language=PostgreSQL
   HelloWorld func() (text string, err error) `
      select 'Hello, world!';`
}
```

Комментарий <code>Dao: generate</code> показывает, что по данной структуре нужно генерировать, комментарий
<code>language=PostgreSQL</code> включает language injection в IDE от JetBrains - GoLang. Это даёт всю мощь расширений
по работе с базами данных — проверка синтаксиса SQL, проверка доступности полей, выполнение SQL прямо из строки в
языке Go — вам не придётся логировать запросы вашей страной ORM, что бы разобраться, что произошло. SQL статичен -
можно использовать оператор explain и заранее отпрофелировать по производительности. Теперь несколько более сложных примеров.

```go
package example

// GoDao: generate
type User struct {
   // language=PostgreSQL
   Init func() (err error) `
        create table "user"(
            "id" serial8 primary key,
            "login" text not null
        );`

   // language=PostgreSQL
   Add func(login string) (id int64, err error) `
      insert into "user" ("login") values ($1::text) returning "id";`
}
```

<code>err error</code> последним возвращаемым параметром обязательно, все функции выполнения запросов под капотом возвращают ошибки.
Входные параметры распаковываются в нумерованные переменные <code>$1</code>. Если вы сталкивались с запуском SQL, то вы,
возможно, знаете про разницу между <code>sql.Exec()</code>, <code>sql.QueryRow()</code> и <code>sql.Query()</code> из
пакета sql. Дело в том, что драйвер не знает, что вернулось при запросе. Он должен освободить соединение с базой данных
для следующего запроса, но только после того, как результат прочитан. Потому при нулевом количестве возвращаемых
значений нужно использовать <code>sql.Exec()</code>, при одной строке - <code>sql.QueryRow()</code>, а при
неопределённом количестве параметров — вычитывать построчно из тела ответа, из объекта <code>sql.Rows</code>, не забыв
вручную закрыть его. Однако это лишнее. Библиотека скрывает сложности интерфейса, от вас требуется только согласовать
тип возвращаемого значения и SQL запрос. Лучше на примере:

<ul><li>

Нет результата в SQL — нет возвращаемой переменной.

```go
package example

// GoDao: generate
type GoDao2 struct {
   // language=PostgreSQL
   DropTestDatabase func() (err error) `
        drop database if exists "test";`
}
```

</li><li>

Возвращается нативный для postgres тип, один или  — в возвращаемых значениях тоже примитивный тип. Таблица нативных типов ниже.

```go
package example

import "github.com/jackc/pgtype"

// GoDao: generate
type GoDao3 struct {
   // language=PostgreSQL
   GetSettings func(id int64) (json pgtype.JSON, err error) `
        with "tmp"("k", "v") as (values
            (0::int8, '{"dark_theme": true}'::json),
            (1::int8, '{"cookies": false}'::json)
        ) select "v"
        from "tmp"
        where "k" = $1
        limit 1;`
}
```

</li><li>

Если вы запрашиваете 1 колонку, то проще и быстрее создать массив из результатов запроса и вернуть его. Все нативные
типы Postgres имеют производный нативный тип массива.

```go
package example

// GoDao: generate
type GoDao5 struct {
   // Расчёт разницы между соседними значениями:
   // [42, 43, 38, 35, 37, 35, 36, 33] => [42, 1, -5, -3, 2, -2, 1, -3, -33]
   // language=PostgreSQL
   Delta func(input []int32) (output []int32, err error) `
        select array(
            select "tmp"."to" - "tmp"."from"
            from unnest(0 || $1::int4[], $1::int4[] || 0) as "tmp"("from", "to")
        )::int4[];`
}
```

</li><li>

Если вы запрашиваете несколько строчек, то используйте массив структур. Порядок возвращаемых значений должен совпадать
с порядком полей в структуре. Кстати, обратите внимание, как порядок сортировки изменяется логической переменной. Любую
логику можно написать, не используя шаблонизацию, SQL достаточно гибок.  

```go
package example

import "github.com/jackc/pgtype"

type Setting struct {
   Key   int64
   Value pgtype.JSON
}

// GoDao: generate
type GoDao7 struct {
   // language=PostgreSQL
   SelectUsers func(ascendingOrder bool, deleted bool) (settings []Setting, err error) `
        with "tmp" ("key", "value") as (values
            (1, '{"name": "Павел Дуров"}'::json),
            (2, '{"name": "Александра Владимирова", "deleted": true}'::json),
            (3, '{"name": "Вячеслав Мирилашвили", "deleted": true}'::json),
            (4, '{"name": "Лев Левиев", "deleted": true}'::json)
            -- отсылка https://ru.wikipedia.org/wiki/Код_Дурова                            
        ) select "key", "value"
        from "tmp"
        where coalesce("value"->>'deleted', false)::bool = $2::bool
        order by
            case when $1::bool then "key" end desc,
            case when not $1::bool then "key" end asc;`
}
```

</li></ul>

<h3>Генерация</h3>

После реализации управляющих структур надо запустить генератор.
Лучше всего использовать встроеный функционал go generate. 
Скажем, в файле users_storage.go добавьте инструкцию

```go
//go:generate go run github.com/OlegSchwann/GoDao ./users_storage.go
```

GoLand сразу предложит запустить генератор маленькой зелёной стрелочкой. Но можно сделать это и из консоли запустив go generate.

<h3>Система контроля версий</h3>

Рекомендуется commit'ить сгенерированный код в репозиторий, вместе с вручную сделанными изменениями.
Это позволит запускать ваш проект сразу после <code>go get</code>, без сложных настроек. 

<h3>Нативные типы PostgreSQL</h3>

Некоторые типы являются нативными для базы данных. Это позволяет использовать
бинарный высокопроизводительный протокол передачи данных.  Вот полный список доступных нативных типов:

<table>
  <thead>
    <tr>
      <td>postgres type with aliases</td>
      <td>golang type</td>
      <td>description</td>
    </tr>
  </thead><tbody>
    <tr>
      <td>int2, smallint, serial2, smallserial</td>
      <td>int16, pgtype.Int2</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-numeric.html#DATATYPE-INT" target="_blank">signed two-byte integer</a></td>
    </tr><tr>
      <td>int2[], smallint[]</td>
      <td>[]int16, []pgtype.Int2, pgtype.Int2Array</td>
      <td>array of two-byte integers</td>
    </tr><tr>
      <td>int4, int, integer, serial4, serial</td>
      <td>int32, pgtype.Int4</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-numeric.html#DATATYPE-INT" target="_blank">signed four-byte integer</a></td>
    </tr><tr>
      <td>int4[], int[], integer[]</td>
      <td>[]int32, []pgtype.Int4, pgtype.Int4Array</td>
      <td>array of four-byte integer</td>
    </tr><tr>
      <td>int8, bigint, serial8, bigserial</td>
      <td>int64, pgtype.Int8</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-numeric.html#DATATYPE-INT" target="_blank">signed eight-byte integer</a></td>
    </tr><tr>
      <td>int8[], bigint[]</td>
      <td>[]int64, []pgtype.Int8, pgtype.Int8Array</td>
      <td>array of eight-byte integers</td>
    </tr><tr>
      <td>numeric, decimal</td>
      <td>pgtype.Numeric</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-numeric.html#DATATYPE-NUMERIC-DECIMAL" target="_blank">exact numeric of selectable precision</a></td>
    </tr><tr>
      <td>numeric[], decimal[]</td>
      <td>[]pgtype.Numeric, pgtype.NumericArray</td>
      <td>array of numeric</td>
    </tr><tr>
      <td>float4, real</td>
      <td>float32, pgtype.Float4</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-numeric.html#DATATYPE-FLOAT" target="_blank">single precision floating-point number (4 bytes)</a></td>
    </tr><tr>
      <td>float4[], real[]</td>
      <td>[]float32, []pgtype.Float4, pgtype.Float4Array</td>
      <td>array of floating-point numbers (4 bytes)</td>
    </tr><tr>
      <td>float8, double precision</td>
      <td>float64, pgtype.Float8</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-numeric.html#DATATYPE-FLOAT" target="_blank">double precision floating-point number (8 bytes)</a></td>
    </tr><tr>
      <td>float8[], double precision[]</td>
      <td>[]float64, []pgtype.Float8, pgtype.Float8Array</td>
      <td>array of floating-point number (8 bytes)</td>
    </tr><tr>
      <td>varchar, character varying</td>
      <td>pgtype.Varchar</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-character.html" target="_blank">variable-length character string</a></td>
    </tr><tr>
      <td>varchar[], character varying[]</td>
      <td>[]pgtype.Varchar, pgtype.VarcharArray</td>
      <td>array of strings with limit length</td>
    </tr><tr>
      <td>bpchar</td>
      <td>pgtype.BPChar</td>
      <td>blank-padded string</td>
    </tr><tr>
      <td>bpchar[]</td>
      <td>[]pgtype.BPChar, pgtype.BPCharArray</td>
      <td>array of blank-padded string</td>
    </tr><tr>
      <td>text</td>
      <td>string, pgtype.Text</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-character.html" target="_blank">variable-length character string</a></td>
    </tr><tr>
      <td>text[]</td>
      <td>[]string, []pgtype.Text, pgtype.TextArray</td>
      <td>array of strings without length limit</td>
    </tr><tr>
      <td>"char"</td>
      <td>int8, byte, pgtype.QChar</td>
      <td>C language char - 8 bit. You can not store many byte characters in it, '字'::"char" is error. Quotes are obligatory, without them it will be another type.</td>
    </tr><tr>
      <td>bytea</td>
      <td>[]byte, []int8, pgtype.Bytea</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-binary.html" target="_blank">binary data (“byte array”)</a></td>
    </tr><tr>
      <td>bytea[]</td>
      <td>[][]byte, [][]int8, []pgtype.Bytea, pgtype.ByteaArray</td>
      <td>array of byte strings</td>
    </tr><tr>
      <td>date</td>
      <td>pgtype.Date</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-datetime.html" target="_blank">calendar date (year, month, day)</a></td>
    </tr><tr>
      <td>date[]</td>
      <td>[]pgtype.Date, pgtype.DateArray</td>
      <td>array of calendar dates</td>
    </tr><tr>
      <td>time, time without time zone</td>
      <td>pgtype.Time</td>
      <td>time of day (no time zone)</td>
    </tr><tr>
      <td>timestamp, timestamp without time zone</td>
      <td>pgtype.Timestamp</td>
      <td>date and time (no time zone)</td>
    </tr><tr>
      <td>timestamp[], timestamp without time zone[]</td>
      <td>[]pgtype.Timestamp, pgtype.TimestampArray</td>
      <td>array of unix timestamps</td>
    </tr><tr>
      <td>timestamptz, timestamp with time zone</td>
      <td>time.Time, pgtype.Timestamptz</td>
      <td>date and time, including time zone</td>
    </tr><tr>
      <td>timestamptz[], timestamp with time zone[]</td>
      <td>[]time.Time, []pgtype.Timestamptz, pgtype.TimestamptzArray</td>
      <td>array of time with timezone</td>
    </tr><tr>
      <td>interval</td>
      <td>time.Duration, pgtype.Interval</td>
      <td>time span</td>
    </tr><tr>
      <td>bool, boolean</td>
      <td>bool, pgtype.Bool</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-boolean.html" target="_blank">logical Boolean (true/false)</a></td>
    </tr><tr>
      <td>bool[], boolean[]</td>
      <td>[]bool, []pgtype.Bool, pgtype.BoolArray</td>
      <td>array of boolean</td>
    </tr><tr>
      <td>point</td>
      <td>pgtype.Point</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-geometric.html#id-1.5.7.16.5" target="_blank">geometric point on a plane</a></td>
    </tr><tr>
      <td>line</td>
      <td>pgtype.Line</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-geometric.html#DATATYPE-LINE" target="_blank">infinite line on a plane</a></td>
    </tr><tr>
      <td>lseg</td>
      <td>pgtype.Lseg</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-geometric.html#DATATYPE-LSEG" target="_blank">line segment on a plane</a></td>
    </tr><tr>
      <td>box</td>
      <td>pgtype.Box</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-geometric.html#id-1.5.7.16.8" target="_blank">rectangular box on a plane</a></td>
    </tr><tr>
      <td>path</td>
      <td>pgtype.Path</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-geometric.html#id-1.5.7.16.9" target="_blank">geometric path on a plane</a></td>
    </tr><tr>
      <td>polygon</td>
      <td>pgtype.Polygon</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-geometric.html#DATATYPE-POLYGON" target="_blank">closed geometric path on a plane</a></td>
    </tr><tr>
      <td>circle</td>
      <td>pgtype.Circle</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-geometric.html#DATATYPE-CIRCLE" target="_blank">circle on a plane</a></td>
    </tr><tr>
      <td>inet</td>
      <td>net.IPNet, pgtype.Inet</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-net-types.html#DATATYPE-INET" target="_blank">IPv4 or IPv6 host address, subnetted</a></td>
    </tr><tr>
      <td>inet[]</td>
      <td>[]net.IPNet, []pgtype.Inet, pgtype.InetArray</td>
      <td>array of IPv4 or IPv6</td>
    </tr><tr>
      <td>cidr</td>
      <td>pgtype.CIDR</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-net-types.html#DATATYPE-CIDR" target="_blank">IPv4 or IPv6 network address, the bits of the device address range in this network must be set to 0</a></td>
    </tr><tr>
      <td>cidr[]</td>
      <td>[]pgtype.CIDR, pgtype.CIDRArray</td>
      <td>array of IPv4 or IPv6 network addresses</td>
    </tr><tr>
      <td>macaddr</td>
      <td>net.HardwareAddr, pgtype.Macaddr</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-net-types.html#DATATYPE-MACADDR" target="_blank">MAC (Media Access Control) address</a></td>
    </tr><tr>
      <td>bit</td>
      <td>pgtype.Bit</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-bit.html" target="_blank">fixed-length bit string</a></td>
    </tr><tr>
      <td>varbit, bit varying</td>
      <td>pgtype.Varbit</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-bit.html" target="_blank">variable-length bit string</a></td>
    </tr><tr>
      <td>uuid</td>
      <td>[16]byte, pgtype.UUID</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-uuid.html" target="_blank">universally unique identifier</a></td>
    </tr><tr>
      <td>uuid[]</td>
      <td>[][16]byte, []pgtype.UUID, pgtype.UUIDArray</td>
      <td>array of uuid</td>
    </tr><tr>
      <td>json</td>
      <td>pgtype.JSON</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-json.html" target="_blank">textual JSON data</a></td>
    </tr><tr>
      <td>jsonb</td>
      <td>pgtype.JSONB</td>
      <td><a href="https://www.postgresql.org/docs/current/datatype-json.html#JSON-CONTAINMENT" target="_blank">binary JSON data, decomposed</a></td>
    </tr><tr>
      <td>int4range</td>
      <td>pgtype.Int4range</td>
      <td><a href="https://www.postgresql.org/docs/current/rangetypes.html" target="_blank">range of integer, can have inclusive and exclusive bounds</a></td>
    </tr><tr>
      <td>int8range</td>
      <td>pgtype.Int8range</td>
      <td><a href="https://www.postgresql.org/docs/current/rangetypes.html" target="_blank">range of bigint</a></td>
    </tr><tr>
      <td>numrange</td>
      <td>pgtype.Numrange</td>
      <td><a href="https://www.postgresql.org/docs/current/rangetypes.html" target="_blank">range of numeric</a></td>
    </tr><tr>
      <td>tsrange</td>
      <td>pgtype.Tsrange</td>
      <td><a href="https://www.postgresql.org/docs/current/rangetypes.html" target="_blank">range of timestamp without time zone</a></td>
    </tr><tr>
      <td>tstzrange</td>
      <td>pgtype.Tstzrange</td>
      <td><a href="https://www.postgresql.org/docs/current/rangetypes.html" target="_blank">range of timestamp with time zone</a></td>
    </tr><tr>
      <td>daterange</td>
      <td>pgtype.Daterange</td>
      <td><a href="https://www.postgresql.org/docs/current/rangetypes.html" target="_blank">Range of date</a></td>
    </tr><tr>
      <td>aclitem</td>
      <td>pgtype.ACLItem</td>
      <td><a href="https://github.com/jackc/pgtype/blob/9e700ff067212a8c0d4a2020825a219f045b7571/aclitem.go#L21" target="_blank">aclitem data type</a></td>
    </tr><tr>
      <td>aclitem[]</td>
      <td>[]pgtype.ACLItem, pgtype.ACLItemArray</td>
      <td>array of ACLItem</td>
    </tr><tr>
      <td>name</td>
      <td>pgtype.Name</td>
      <td><a href="https://github.com/jackc/pgtype/blob/9e700ff067212a8c0d4a2020825a219f045b7571/name.go#L20" target="_blank">entity name in postgres</a></td>
    </tr><tr>
      <td>oid</td>
      <td>pgtype.OID</td>
      <td><a href="https://github.com/jackc/pgtype/blob/9e700ff067212a8c0d4a2020825a219f045b7571/oid.go#L19" target="_blank">Object Identifier Type</a></td>
    </tr><tr>
      <td>tid</td>
      <td>pgtype.TID</td>
      <td><a href="https://github.com/jackc/pgtype/blob/9e700ff067212a8c0d4a2020825a219f045b7571/tid.go#L25" target="_blank">Tuple Identifier type</a></td>
    </tr><tr>
      <td>xid</td>
      <td>pgtype.XID</td>
      <td><a href="https://github.com/jackc/pgtype/blob/9e700ff067212a8c0d4a2020825a219f045b7571/xid.go#L21" target="_blank">Transaction ID type</a></td>
    </tr><tr>
      <td>cid</td>
      <td>pgtype.CID</td>
      <td><a href="https://github.com/jackc/pgtype/blob/9e700ff067212a8c0d4a2020825a219f045b7571/cid.go#L18" target="_blank">Command Identifier type</a></td>
    </tr>
  </tbody>
</table>

<h3>Вопросы</h3>

Если при использовании возникли вопросы, я буду рад ответить в
<a href="https://github.com/OlegSchwann/GoDao/issues">issues</a> или в 
<a href="https://t.me/MaxDev">Telegram</a>.
