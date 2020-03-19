<h3>Supported types convertions</h3>
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
