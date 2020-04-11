package main

import (
	"bufio"
	"github.com/OlegSchwann/GoDao/internal/ast"
	"github.com/OlegSchwann/GoDao/internal/extract"
	"github.com/OlegSchwann/GoDao/internal/flag"
	"github.com/OlegSchwann/GoDao/internal/template"
	"log"
	"os"
)

// Пишем кодогенератор.
// Он будет работать через // go generate, преимущественно.
// Нужно написать версию 0.1 , только бы работало.

// Этапы работы кодогенератора:
// парсим аргументы командной строки +
// парсим файл
// находим нужные структуры
// находим в структурах нужные структуры
// формируем правило в виде структуры
// в том числе тип
// записываем в файлик с похожим расширением

// начнём с написания тестов.
// пример использования и результата, из которого можно потом наделать шаблонов.

func main() {
	config, err := flag.ShellParsing(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	file, err := ast.ParseFile(config)
	if err != nil {
		log.Fatal(err)
	}

	dot, err := extract.Extract(file)
	if err != nil {
		log.Fatal(err)
	}

	outputFile, err := os.Create(config.OutputGoFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()
	outputFileBuffered := bufio.NewWriter(outputFile)
	defer outputFileBuffered.Flush()
	err = template.Render(dot, outputFileBuffered)
	if err != nil {
		log.Fatal(err)
	}

}

// задача:
// объяснить, как использовать нативные типы postgres через библиотеку.
// задача 2
// используя регулярку \b\$\d+::(\w+)\b для проверки типов во время шаблонизации.
// задача - исчерпывающее соответствие типов go и postgres.

// типы go ---> pgtype типы -+-> типы postges

// замечания
// в библиотеке pgx потерян тип time with timezone
// в библиотеке потерян тип char и character, который строка с фиксированной длинной, обрезающаяся при превышении без предупреждения.
// в библиотеке потерян тип xml
// в документации https://www.postgresql.org/docs/current/datatype.html потерян тип "char", который byte, есть только char, который n симполов.
// numeric и decimal не имеют конструктора из math.big,  не смотря на то, что используют его под капотом.
// в библиотеке отсутствует тип macaddr8
// в библиотеке отсутствует тип macaddr8[]
// в библиотеке отсутствует тип macaddr[]

//TODO добавить Usage: gogao [OPTION] TARGET_FILE

// структура кодогенератора:
// распарсили флаги
// распарсили файл в AST
// распарсили AST в структуру для щаблона
// открыли файл выходной
// шаблонизировали
