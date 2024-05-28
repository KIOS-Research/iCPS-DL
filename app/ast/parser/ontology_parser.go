// Code generated from .\ast\parser\ontology.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ontology

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ontologyParser struct {
	*antlr.BaseParser
}

var ontologyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ontologyParserInit() {
	staticData := &ontologyParserStaticData
	staticData.literalNames = []string{
		"", "'load'", "'show'", "'ls'", "'translate'", "'mermaid'", "'rm'",
		"'traverse'", "'configure'", "'compose'", "'project'", "'info'", "'render'",
		"'exit'", "'quit'", "'paradigm'", "'process'", "'knowledge'", "'base'",
		"'property'", "'model'", "'class'", "'physical'", "'actuator'", "'sensor'",
		"'translation'", "'device'", "'conn'", "'estimate'", "'sense'", "'control'",
		"'actuate'", "'using'", "'local'", "'global'", "'or'", "'end'", "'('",
		"')'", "'{'", "'}'", "'['", "']'", "'@'", "'='", "':='", "','", "':'",
		"'.'", "'->'", "'!'", "'?'",
	}
	staticData.symbolicNames = []string{
		"", "LOAD", "SHOW", "LS", "TRANSLATE", "MERMAID", "REMOVE", "TRAVERSE",
		"CONFIGURE", "COMPOSE", "PROJECT", "INFO", "RENDER", "EXIT", "QUIT",
		"PARADIGM", "PROCESS", "KNOWLEDGE", "BASE", "PROPERTY", "MODEL", "CLASS",
		"PHYSICAL", "ACTUATOR", "SENSOR", "TRANSLATION", "DEVICE", "CONNECTION",
		"ESTIMATE", "SENSE", "CONTROL", "ACTUATE", "USING", "LOCAL", "GLOBAL",
		"OR", "END", "LPAR", "RPAR", "LBRA", "RBRA", "LSQ", "RSQ", "AT", "EQ",
		"ASGN", "COMMA", "COLON", "DOT", "ARROW", "BANG", "QUESTION", "PATH",
		"ID", "INT", "COMMENTS", "WS",
	}
	staticData.ruleNames = []string{
		"program", "exit", "expression", "load", "assignment", "remove", "show",
		"mermaid", "translate", "traverse", "configure", "compose", "project",
		"info", "render", "paradigm", "paradigm_declaration", "property", "type",
		"model", "class", "internal_edge", "translation", "arg_connection",
		"connection", "process", "process_declaration", "device", "component",
		"connection_decl", "local_configuration", "global_configuration", "local",
		"send", "receive", "loop", "label", "select", "branch", "global", "pass",
		"global_loop", "choice", "knowledge_base", "knowledge_base_decl",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 56, 515, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 1, 0, 5, 0, 92, 8, 0, 10, 0, 12,
		0, 95, 9, 0, 1, 0, 3, 0, 98, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 127, 8, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 133, 8, 2, 10, 2, 12, 2, 136, 9, 2, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6,
		150, 8, 6, 1, 6, 1, 6, 3, 6, 154, 8, 6, 3, 6, 156, 8, 6, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 3, 13, 184, 8, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		5, 15, 192, 8, 15, 10, 15, 12, 15, 195, 9, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 3, 16, 203, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17,
		209, 8, 17, 10, 17, 12, 17, 212, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 5, 18, 219, 8, 18, 10, 18, 12, 18, 222, 9, 18, 1, 18, 1, 18, 3, 18,
		226, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 232, 8, 19, 10, 19, 12,
		19, 235, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 243, 8,
		20, 10, 20, 12, 20, 246, 9, 20, 3, 20, 248, 8, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 268, 8, 22, 10, 22, 12, 22, 271,
		9, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 289, 8, 25, 10, 25,
		12, 25, 292, 9, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 3, 26, 299, 8, 26,
		1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 305, 8, 27, 10, 27, 12, 27, 308, 9,
		27, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 314, 8, 28, 10, 28, 12, 28, 317,
		9, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5,
		28, 328, 8, 28, 10, 28, 12, 28, 331, 9, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 342, 8, 28, 10, 28, 12, 28, 345,
		9, 28, 1, 28, 3, 28, 348, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 354,
		8, 29, 10, 29, 12, 29, 357, 9, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 4,
		30, 364, 8, 30, 11, 30, 12, 30, 365, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 380, 8, 32, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 4, 37, 415, 8, 37, 11, 37, 12, 37, 416, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 4, 38, 434, 8, 38, 11, 38, 12, 38, 435, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 3, 39, 443, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 42, 1, 42, 4, 42, 474, 8, 42, 11, 42, 12, 42, 475, 1, 43, 1, 43, 1,
		43, 1, 43, 1, 43, 4, 43, 483, 8, 43, 11, 43, 12, 43, 484, 1, 43, 1, 43,
		1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1,
		44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44,
		1, 44, 1, 44, 1, 44, 3, 44, 513, 8, 44, 1, 44, 0, 1, 4, 45, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
		80, 82, 84, 86, 88, 0, 3, 1, 0, 13, 14, 1, 0, 52, 53, 1, 0, 22, 23, 533,
		0, 93, 1, 0, 0, 0, 2, 101, 1, 0, 0, 0, 4, 126, 1, 0, 0, 0, 6, 137, 1, 0,
		0, 0, 8, 140, 1, 0, 0, 0, 10, 144, 1, 0, 0, 0, 12, 155, 1, 0, 0, 0, 14,
		157, 1, 0, 0, 0, 16, 160, 1, 0, 0, 0, 18, 163, 1, 0, 0, 0, 20, 169, 1,
		0, 0, 0, 22, 175, 1, 0, 0, 0, 24, 178, 1, 0, 0, 0, 26, 181, 1, 0, 0, 0,
		28, 185, 1, 0, 0, 0, 30, 188, 1, 0, 0, 0, 32, 202, 1, 0, 0, 0, 34, 204,
		1, 0, 0, 0, 36, 225, 1, 0, 0, 0, 38, 227, 1, 0, 0, 0, 40, 236, 1, 0, 0,
		0, 42, 255, 1, 0, 0, 0, 44, 259, 1, 0, 0, 0, 46, 272, 1, 0, 0, 0, 48, 280,
		1, 0, 0, 0, 50, 284, 1, 0, 0, 0, 52, 298, 1, 0, 0, 0, 54, 300, 1, 0, 0,
		0, 56, 347, 1, 0, 0, 0, 58, 349, 1, 0, 0, 0, 60, 358, 1, 0, 0, 0, 62, 369,
		1, 0, 0, 0, 64, 379, 1, 0, 0, 0, 66, 381, 1, 0, 0, 0, 68, 387, 1, 0, 0,
		0, 70, 393, 1, 0, 0, 0, 72, 397, 1, 0, 0, 0, 74, 399, 1, 0, 0, 0, 76, 418,
		1, 0, 0, 0, 78, 442, 1, 0, 0, 0, 80, 444, 1, 0, 0, 0, 82, 452, 1, 0, 0,
		0, 84, 456, 1, 0, 0, 0, 86, 477, 1, 0, 0, 0, 88, 512, 1, 0, 0, 0, 90, 92,
		3, 4, 2, 0, 91, 90, 1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0,
		93, 94, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 96, 98, 3,
		2, 1, 0, 97, 96, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99,
		100, 5, 0, 0, 1, 100, 1, 1, 0, 0, 0, 101, 102, 7, 0, 0, 0, 102, 3, 1, 0,
		0, 0, 103, 104, 6, 2, -1, 0, 104, 127, 3, 30, 15, 0, 105, 127, 3, 50, 25,
		0, 106, 127, 3, 60, 30, 0, 107, 127, 3, 62, 31, 0, 108, 127, 3, 86, 43,
		0, 109, 127, 3, 6, 3, 0, 110, 127, 3, 8, 4, 0, 111, 127, 3, 10, 5, 0, 112,
		127, 3, 12, 6, 0, 113, 127, 3, 14, 7, 0, 114, 127, 3, 16, 8, 0, 115, 127,
		3, 18, 9, 0, 116, 127, 3, 20, 10, 0, 117, 127, 3, 22, 11, 0, 118, 127,
		3, 24, 12, 0, 119, 127, 3, 26, 13, 0, 120, 127, 3, 28, 14, 0, 121, 127,
		5, 53, 0, 0, 122, 123, 5, 37, 0, 0, 123, 124, 3, 4, 2, 0, 124, 125, 5,
		38, 0, 0, 125, 127, 1, 0, 0, 0, 126, 103, 1, 0, 0, 0, 126, 105, 1, 0, 0,
		0, 126, 106, 1, 0, 0, 0, 126, 107, 1, 0, 0, 0, 126, 108, 1, 0, 0, 0, 126,
		109, 1, 0, 0, 0, 126, 110, 1, 0, 0, 0, 126, 111, 1, 0, 0, 0, 126, 112,
		1, 0, 0, 0, 126, 113, 1, 0, 0, 0, 126, 114, 1, 0, 0, 0, 126, 115, 1, 0,
		0, 0, 126, 116, 1, 0, 0, 0, 126, 117, 1, 0, 0, 0, 126, 118, 1, 0, 0, 0,
		126, 119, 1, 0, 0, 0, 126, 120, 1, 0, 0, 0, 126, 121, 1, 0, 0, 0, 126,
		122, 1, 0, 0, 0, 127, 134, 1, 0, 0, 0, 128, 129, 10, 3, 0, 0, 129, 130,
		5, 41, 0, 0, 130, 131, 5, 54, 0, 0, 131, 133, 5, 42, 0, 0, 132, 128, 1,
		0, 0, 0, 133, 136, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0,
		0, 135, 5, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 137, 138, 5, 1, 0, 0, 138,
		139, 7, 1, 0, 0, 139, 7, 1, 0, 0, 0, 140, 141, 5, 53, 0, 0, 141, 142, 5,
		45, 0, 0, 142, 143, 3, 4, 2, 0, 143, 9, 1, 0, 0, 0, 144, 145, 5, 6, 0,
		0, 145, 146, 5, 53, 0, 0, 146, 11, 1, 0, 0, 0, 147, 149, 5, 2, 0, 0, 148,
		150, 3, 4, 2, 0, 149, 148, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 156,
		1, 0, 0, 0, 151, 153, 5, 3, 0, 0, 152, 154, 3, 4, 2, 0, 153, 152, 1, 0,
		0, 0, 153, 154, 1, 0, 0, 0, 154, 156, 1, 0, 0, 0, 155, 147, 1, 0, 0, 0,
		155, 151, 1, 0, 0, 0, 156, 13, 1, 0, 0, 0, 157, 158, 5, 5, 0, 0, 158, 159,
		3, 4, 2, 0, 159, 15, 1, 0, 0, 0, 160, 161, 5, 4, 0, 0, 161, 162, 3, 4,
		2, 0, 162, 17, 1, 0, 0, 0, 163, 164, 5, 7, 0, 0, 164, 165, 5, 53, 0, 0,
		165, 166, 5, 48, 0, 0, 166, 167, 5, 53, 0, 0, 167, 168, 3, 4, 2, 0, 168,
		19, 1, 0, 0, 0, 169, 170, 5, 8, 0, 0, 170, 171, 3, 4, 2, 0, 171, 172, 3,
		4, 2, 0, 172, 173, 5, 53, 0, 0, 173, 174, 5, 53, 0, 0, 174, 21, 1, 0, 0,
		0, 175, 176, 5, 9, 0, 0, 176, 177, 3, 4, 2, 0, 177, 23, 1, 0, 0, 0, 178,
		179, 5, 10, 0, 0, 179, 180, 3, 4, 2, 0, 180, 25, 1, 0, 0, 0, 181, 183,
		5, 11, 0, 0, 182, 184, 3, 4, 2, 0, 183, 182, 1, 0, 0, 0, 183, 184, 1, 0,
		0, 0, 184, 27, 1, 0, 0, 0, 185, 186, 5, 12, 0, 0, 186, 187, 3, 4, 2, 0,
		187, 29, 1, 0, 0, 0, 188, 189, 5, 15, 0, 0, 189, 193, 5, 39, 0, 0, 190,
		192, 3, 32, 16, 0, 191, 190, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191,
		1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 196, 1, 0, 0, 0, 195, 193, 1, 0,
		0, 0, 196, 197, 5, 40, 0, 0, 197, 31, 1, 0, 0, 0, 198, 203, 3, 34, 17,
		0, 199, 203, 3, 40, 20, 0, 200, 203, 3, 38, 19, 0, 201, 203, 3, 44, 22,
		0, 202, 198, 1, 0, 0, 0, 202, 199, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 202,
		201, 1, 0, 0, 0, 203, 33, 1, 0, 0, 0, 204, 205, 5, 19, 0, 0, 205, 210,
		3, 36, 18, 0, 206, 207, 5, 46, 0, 0, 207, 209, 3, 36, 18, 0, 208, 206,
		1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 211, 1, 0,
		0, 0, 211, 35, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 213, 214, 5, 53, 0, 0,
		214, 215, 5, 39, 0, 0, 215, 220, 5, 53, 0, 0, 216, 217, 5, 46, 0, 0, 217,
		219, 5, 53, 0, 0, 218, 216, 1, 0, 0, 0, 219, 222, 1, 0, 0, 0, 220, 218,
		1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 223, 1, 0, 0, 0, 222, 220, 1, 0,
		0, 0, 223, 226, 5, 40, 0, 0, 224, 226, 5, 53, 0, 0, 225, 213, 1, 0, 0,
		0, 225, 224, 1, 0, 0, 0, 226, 37, 1, 0, 0, 0, 227, 228, 5, 20, 0, 0, 228,
		233, 5, 53, 0, 0, 229, 230, 5, 46, 0, 0, 230, 232, 5, 53, 0, 0, 231, 229,
		1, 0, 0, 0, 232, 235, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 233, 234, 1, 0,
		0, 0, 234, 39, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 236, 237, 7, 2, 0, 0,
		237, 238, 5, 53, 0, 0, 238, 247, 5, 37, 0, 0, 239, 244, 5, 53, 0, 0, 240,
		241, 5, 46, 0, 0, 241, 243, 5, 53, 0, 0, 242, 240, 1, 0, 0, 0, 243, 246,
		1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 248, 1, 0,
		0, 0, 246, 244, 1, 0, 0, 0, 247, 239, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0,
		248, 249, 1, 0, 0, 0, 249, 250, 5, 38, 0, 0, 250, 251, 5, 47, 0, 0, 251,
		252, 3, 42, 21, 0, 252, 253, 5, 46, 0, 0, 253, 254, 3, 42, 21, 0, 254,
		41, 1, 0, 0, 0, 255, 256, 5, 53, 0, 0, 256, 257, 5, 49, 0, 0, 257, 258,
		5, 53, 0, 0, 258, 43, 1, 0, 0, 0, 259, 260, 5, 25, 0, 0, 260, 261, 5, 53,
		0, 0, 261, 262, 5, 49, 0, 0, 262, 263, 5, 53, 0, 0, 263, 264, 5, 47, 0,
		0, 264, 269, 3, 46, 23, 0, 265, 266, 5, 46, 0, 0, 266, 268, 3, 46, 23,
		0, 267, 265, 1, 0, 0, 0, 268, 271, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 269,
		270, 1, 0, 0, 0, 270, 45, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 272, 273, 5,
		53, 0, 0, 273, 274, 5, 48, 0, 0, 274, 275, 5, 53, 0, 0, 275, 276, 5, 49,
		0, 0, 276, 277, 5, 53, 0, 0, 277, 278, 5, 48, 0, 0, 278, 279, 5, 53, 0,
		0, 279, 47, 1, 0, 0, 0, 280, 281, 5, 53, 0, 0, 281, 282, 5, 49, 0, 0, 282,
		283, 5, 53, 0, 0, 283, 49, 1, 0, 0, 0, 284, 285, 5, 16, 0, 0, 285, 286,
		5, 53, 0, 0, 286, 290, 5, 39, 0, 0, 287, 289, 3, 52, 26, 0, 288, 287, 1,
		0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 291, 1, 0, 0,
		0, 291, 293, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 293, 294, 5, 40, 0, 0, 294,
		51, 1, 0, 0, 0, 295, 299, 3, 54, 27, 0, 296, 299, 3, 56, 28, 0, 297, 299,
		3, 58, 29, 0, 298, 295, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 297, 1,
		0, 0, 0, 299, 53, 1, 0, 0, 0, 300, 301, 5, 26, 0, 0, 301, 306, 5, 53, 0,
		0, 302, 303, 5, 46, 0, 0, 303, 305, 5, 53, 0, 0, 304, 302, 1, 0, 0, 0,
		305, 308, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307,
		55, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 309, 310, 5, 22, 0, 0, 310, 315,
		5, 53, 0, 0, 311, 312, 5, 46, 0, 0, 312, 314, 5, 53, 0, 0, 313, 311, 1,
		0, 0, 0, 314, 317, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 315, 316, 1, 0, 0,
		0, 316, 318, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 318, 348, 5, 53, 0, 0, 319,
		320, 5, 23, 0, 0, 320, 321, 5, 53, 0, 0, 321, 322, 5, 43, 0, 0, 322, 329,
		5, 53, 0, 0, 323, 324, 5, 46, 0, 0, 324, 325, 5, 53, 0, 0, 325, 326, 5,
		43, 0, 0, 326, 328, 5, 53, 0, 0, 327, 323, 1, 0, 0, 0, 328, 331, 1, 0,
		0, 0, 329, 327, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330, 332, 1, 0, 0, 0,
		331, 329, 1, 0, 0, 0, 332, 348, 5, 53, 0, 0, 333, 334, 5, 24, 0, 0, 334,
		335, 5, 53, 0, 0, 335, 336, 5, 43, 0, 0, 336, 343, 5, 53, 0, 0, 337, 338,
		5, 46, 0, 0, 338, 339, 5, 53, 0, 0, 339, 340, 5, 43, 0, 0, 340, 342, 5,
		53, 0, 0, 341, 337, 1, 0, 0, 0, 342, 345, 1, 0, 0, 0, 343, 341, 1, 0, 0,
		0, 343, 344, 1, 0, 0, 0, 344, 346, 1, 0, 0, 0, 345, 343, 1, 0, 0, 0, 346,
		348, 5, 53, 0, 0, 347, 309, 1, 0, 0, 0, 347, 319, 1, 0, 0, 0, 347, 333,
		1, 0, 0, 0, 348, 57, 1, 0, 0, 0, 349, 350, 5, 27, 0, 0, 350, 355, 3, 48,
		24, 0, 351, 352, 5, 46, 0, 0, 352, 354, 3, 48, 24, 0, 353, 351, 1, 0, 0,
		0, 354, 357, 1, 0, 0, 0, 355, 353, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356,
		59, 1, 0, 0, 0, 357, 355, 1, 0, 0, 0, 358, 359, 5, 33, 0, 0, 359, 363,
		5, 39, 0, 0, 360, 361, 5, 53, 0, 0, 361, 362, 5, 44, 0, 0, 362, 364, 3,
		64, 32, 0, 363, 360, 1, 0, 0, 0, 364, 365, 1, 0, 0, 0, 365, 363, 1, 0,
		0, 0, 365, 366, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 368, 5, 40, 0, 0,
		368, 61, 1, 0, 0, 0, 369, 370, 5, 34, 0, 0, 370, 371, 3, 78, 39, 0, 371,
		63, 1, 0, 0, 0, 372, 380, 5, 36, 0, 0, 373, 380, 3, 66, 33, 0, 374, 380,
		3, 68, 34, 0, 375, 380, 3, 70, 35, 0, 376, 380, 3, 72, 36, 0, 377, 380,
		3, 74, 37, 0, 378, 380, 3, 76, 38, 0, 379, 372, 1, 0, 0, 0, 379, 373, 1,
		0, 0, 0, 379, 374, 1, 0, 0, 0, 379, 375, 1, 0, 0, 0, 379, 376, 1, 0, 0,
		0, 379, 377, 1, 0, 0, 0, 379, 378, 1, 0, 0, 0, 380, 65, 1, 0, 0, 0, 381,
		382, 5, 53, 0, 0, 382, 383, 5, 50, 0, 0, 383, 384, 3, 36, 18, 0, 384, 385,
		5, 48, 0, 0, 385, 386, 3, 64, 32, 0, 386, 67, 1, 0, 0, 0, 387, 388, 5,
		53, 0, 0, 388, 389, 5, 51, 0, 0, 389, 390, 3, 36, 18, 0, 390, 391, 5, 48,
		0, 0, 391, 392, 3, 64, 32, 0, 392, 69, 1, 0, 0, 0, 393, 394, 5, 53, 0,
		0, 394, 395, 5, 48, 0, 0, 395, 396, 3, 64, 32, 0, 396, 71, 1, 0, 0, 0,
		397, 398, 5, 53, 0, 0, 398, 73, 1, 0, 0, 0, 399, 400, 5, 53, 0, 0, 400,
		401, 5, 50, 0, 0, 401, 402, 5, 53, 0, 0, 402, 403, 5, 39, 0, 0, 403, 404,
		5, 53, 0, 0, 404, 405, 5, 47, 0, 0, 405, 406, 3, 64, 32, 0, 406, 414, 5,
		40, 0, 0, 407, 408, 5, 35, 0, 0, 408, 409, 5, 39, 0, 0, 409, 410, 5, 53,
		0, 0, 410, 411, 5, 47, 0, 0, 411, 412, 3, 64, 32, 0, 412, 413, 5, 40, 0,
		0, 413, 415, 1, 0, 0, 0, 414, 407, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0, 416,
		414, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 75, 1, 0, 0, 0, 418, 419, 5,
		53, 0, 0, 419, 420, 5, 51, 0, 0, 420, 421, 5, 53, 0, 0, 421, 422, 5, 39,
		0, 0, 422, 423, 5, 53, 0, 0, 423, 424, 5, 47, 0, 0, 424, 425, 3, 64, 32,
		0, 425, 433, 5, 40, 0, 0, 426, 427, 5, 35, 0, 0, 427, 428, 5, 39, 0, 0,
		428, 429, 5, 53, 0, 0, 429, 430, 5, 47, 0, 0, 430, 431, 3, 64, 32, 0, 431,
		432, 5, 40, 0, 0, 432, 434, 1, 0, 0, 0, 433, 426, 1, 0, 0, 0, 434, 435,
		1, 0, 0, 0, 435, 433, 1, 0, 0, 0, 435, 436, 1, 0, 0, 0, 436, 77, 1, 0,
		0, 0, 437, 443, 3, 80, 40, 0, 438, 443, 3, 82, 41, 0, 439, 443, 3, 72,
		36, 0, 440, 443, 3, 84, 42, 0, 441, 443, 5, 36, 0, 0, 442, 437, 1, 0, 0,
		0, 442, 438, 1, 0, 0, 0, 442, 439, 1, 0, 0, 0, 442, 440, 1, 0, 0, 0, 442,
		441, 1, 0, 0, 0, 443, 79, 1, 0, 0, 0, 444, 445, 5, 53, 0, 0, 445, 446,
		5, 49, 0, 0, 446, 447, 5, 53, 0, 0, 447, 448, 5, 47, 0, 0, 448, 449, 3,
		36, 18, 0, 449, 450, 5, 48, 0, 0, 450, 451, 3, 78, 39, 0, 451, 81, 1, 0,
		0, 0, 452, 453, 5, 53, 0, 0, 453, 454, 5, 48, 0, 0, 454, 455, 3, 78, 39,
		0, 455, 83, 1, 0, 0, 0, 456, 457, 5, 53, 0, 0, 457, 458, 5, 49, 0, 0, 458,
		459, 5, 53, 0, 0, 459, 460, 5, 47, 0, 0, 460, 461, 5, 53, 0, 0, 461, 462,
		5, 39, 0, 0, 462, 463, 5, 53, 0, 0, 463, 464, 5, 47, 0, 0, 464, 465, 3,
		78, 39, 0, 465, 473, 5, 40, 0, 0, 466, 467, 5, 35, 0, 0, 467, 468, 5, 39,
		0, 0, 468, 469, 5, 53, 0, 0, 469, 470, 5, 47, 0, 0, 470, 471, 3, 78, 39,
		0, 471, 472, 5, 40, 0, 0, 472, 474, 1, 0, 0, 0, 473, 466, 1, 0, 0, 0, 474,
		475, 1, 0, 0, 0, 475, 473, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476, 85, 1,
		0, 0, 0, 477, 478, 5, 17, 0, 0, 478, 479, 5, 18, 0, 0, 479, 480, 5, 53,
		0, 0, 480, 482, 5, 39, 0, 0, 481, 483, 3, 88, 44, 0, 482, 481, 1, 0, 0,
		0, 483, 484, 1, 0, 0, 0, 484, 482, 1, 0, 0, 0, 484, 485, 1, 0, 0, 0, 485,
		486, 1, 0, 0, 0, 486, 487, 5, 40, 0, 0, 487, 87, 1, 0, 0, 0, 488, 489,
		5, 28, 0, 0, 489, 490, 5, 53, 0, 0, 490, 491, 5, 32, 0, 0, 491, 492, 5,
		53, 0, 0, 492, 493, 5, 44, 0, 0, 493, 513, 3, 64, 32, 0, 494, 495, 5, 29,
		0, 0, 495, 496, 5, 53, 0, 0, 496, 497, 5, 32, 0, 0, 497, 498, 5, 53, 0,
		0, 498, 499, 5, 44, 0, 0, 499, 513, 3, 64, 32, 0, 500, 501, 5, 30, 0, 0,
		501, 502, 5, 53, 0, 0, 502, 503, 5, 32, 0, 0, 503, 504, 5, 53, 0, 0, 504,
		505, 5, 44, 0, 0, 505, 513, 3, 64, 32, 0, 506, 507, 5, 31, 0, 0, 507, 508,
		5, 53, 0, 0, 508, 509, 5, 32, 0, 0, 509, 510, 5, 53, 0, 0, 510, 511, 5,
		44, 0, 0, 511, 513, 3, 64, 32, 0, 512, 488, 1, 0, 0, 0, 512, 494, 1, 0,
		0, 0, 512, 500, 1, 0, 0, 0, 512, 506, 1, 0, 0, 0, 513, 89, 1, 0, 0, 0,
		33, 93, 97, 126, 134, 149, 153, 155, 183, 193, 202, 210, 220, 225, 233,
		244, 247, 269, 290, 298, 306, 315, 329, 343, 347, 355, 365, 379, 416, 435,
		442, 475, 484, 512,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ontologyParserInit initializes any static state used to implement ontologyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewontologyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OntologyParserInit() {
	staticData := &ontologyParserStaticData
	staticData.once.Do(ontologyParserInit)
}

// NewontologyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewontologyParser(input antlr.TokenStream) *ontologyParser {
	OntologyParserInit()
	this := new(ontologyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ontologyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ontology.g4"

	return this
}

// ontologyParser tokens.
const (
	ontologyParserEOF         = antlr.TokenEOF
	ontologyParserLOAD        = 1
	ontologyParserSHOW        = 2
	ontologyParserLS          = 3
	ontologyParserTRANSLATE   = 4
	ontologyParserMERMAID     = 5
	ontologyParserREMOVE      = 6
	ontologyParserTRAVERSE    = 7
	ontologyParserCONFIGURE   = 8
	ontologyParserCOMPOSE     = 9
	ontologyParserPROJECT     = 10
	ontologyParserINFO        = 11
	ontologyParserRENDER      = 12
	ontologyParserEXIT        = 13
	ontologyParserQUIT        = 14
	ontologyParserPARADIGM    = 15
	ontologyParserPROCESS     = 16
	ontologyParserKNOWLEDGE   = 17
	ontologyParserBASE        = 18
	ontologyParserPROPERTY    = 19
	ontologyParserMODEL       = 20
	ontologyParserCLASS       = 21
	ontologyParserPHYSICAL    = 22
	ontologyParserACTUATOR    = 23
	ontologyParserSENSOR      = 24
	ontologyParserTRANSLATION = 25
	ontologyParserDEVICE      = 26
	ontologyParserCONNECTION  = 27
	ontologyParserESTIMATE    = 28
	ontologyParserSENSE       = 29
	ontologyParserCONTROL     = 30
	ontologyParserACTUATE     = 31
	ontologyParserUSING       = 32
	ontologyParserLOCAL       = 33
	ontologyParserGLOBAL      = 34
	ontologyParserOR          = 35
	ontologyParserEND         = 36
	ontologyParserLPAR        = 37
	ontologyParserRPAR        = 38
	ontologyParserLBRA        = 39
	ontologyParserRBRA        = 40
	ontologyParserLSQ         = 41
	ontologyParserRSQ         = 42
	ontologyParserAT          = 43
	ontologyParserEQ          = 44
	ontologyParserASGN        = 45
	ontologyParserCOMMA       = 46
	ontologyParserCOLON       = 47
	ontologyParserDOT         = 48
	ontologyParserARROW       = 49
	ontologyParserBANG        = 50
	ontologyParserQUESTION    = 51
	ontologyParserPATH        = 52
	ontologyParserID          = 53
	ontologyParserINT         = 54
	ontologyParserCOMMENTS    = 55
	ontologyParserWS          = 56
)

// ontologyParser rules.
const (
	ontologyParserRULE_program              = 0
	ontologyParserRULE_exit                 = 1
	ontologyParserRULE_expression           = 2
	ontologyParserRULE_load                 = 3
	ontologyParserRULE_assignment           = 4
	ontologyParserRULE_remove               = 5
	ontologyParserRULE_show                 = 6
	ontologyParserRULE_mermaid              = 7
	ontologyParserRULE_translate            = 8
	ontologyParserRULE_traverse             = 9
	ontologyParserRULE_configure            = 10
	ontologyParserRULE_compose              = 11
	ontologyParserRULE_project              = 12
	ontologyParserRULE_info                 = 13
	ontologyParserRULE_render               = 14
	ontologyParserRULE_paradigm             = 15
	ontologyParserRULE_paradigm_declaration = 16
	ontologyParserRULE_property             = 17
	ontologyParserRULE_type                 = 18
	ontologyParserRULE_model                = 19
	ontologyParserRULE_class                = 20
	ontologyParserRULE_internal_edge        = 21
	ontologyParserRULE_translation          = 22
	ontologyParserRULE_arg_connection       = 23
	ontologyParserRULE_connection           = 24
	ontologyParserRULE_process              = 25
	ontologyParserRULE_process_declaration  = 26
	ontologyParserRULE_device               = 27
	ontologyParserRULE_component            = 28
	ontologyParserRULE_connection_decl      = 29
	ontologyParserRULE_local_configuration  = 30
	ontologyParserRULE_global_configuration = 31
	ontologyParserRULE_local                = 32
	ontologyParserRULE_send                 = 33
	ontologyParserRULE_receive              = 34
	ontologyParserRULE_loop                 = 35
	ontologyParserRULE_label                = 36
	ontologyParserRULE_select               = 37
	ontologyParserRULE_branch               = 38
	ontologyParserRULE_global               = 39
	ontologyParserRULE_pass                 = 40
	ontologyParserRULE_global_loop          = 41
	ontologyParserRULE_choice               = 42
	ontologyParserRULE_knowledge_base       = 43
	ontologyParserRULE_knowledge_base_decl  = 44
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(ontologyParserEOF, 0)
}

func (s *ProgramContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProgramContext) Exit() IExitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExitContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ontologyParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ontologyParserLOAD)|(1<<ontologyParserSHOW)|(1<<ontologyParserLS)|(1<<ontologyParserTRANSLATE)|(1<<ontologyParserMERMAID)|(1<<ontologyParserREMOVE)|(1<<ontologyParserTRAVERSE)|(1<<ontologyParserCONFIGURE)|(1<<ontologyParserCOMPOSE)|(1<<ontologyParserPROJECT)|(1<<ontologyParserINFO)|(1<<ontologyParserRENDER)|(1<<ontologyParserPARADIGM)|(1<<ontologyParserPROCESS)|(1<<ontologyParserKNOWLEDGE))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(ontologyParserLOCAL-33))|(1<<(ontologyParserGLOBAL-33))|(1<<(ontologyParserLPAR-33))|(1<<(ontologyParserID-33)))) != 0) {
		{
			p.SetState(90)
			p.expression(0)
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ontologyParserEXIT || _la == ontologyParserQUIT {
		{
			p.SetState(96)
			p.Exit()
		}

	}
	{
		p.SetState(99)
		p.Match(ontologyParserEOF)
	}

	return localctx
}

// IExitContext is an interface to support dynamic dispatch.
type IExitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExitContext differentiates from other interfaces.
	IsExitContext()
}

type ExitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExitContext() *ExitContext {
	var p = new(ExitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_exit
	return p
}

func (*ExitContext) IsExitContext() {}

func NewExitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExitContext {
	var p = new(ExitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_exit

	return p
}

func (s *ExitContext) GetParser() antlr.Parser { return s.parser }

func (s *ExitContext) EXIT() antlr.TerminalNode {
	return s.GetToken(ontologyParserEXIT, 0)
}

func (s *ExitContext) QUIT() antlr.TerminalNode {
	return s.GetToken(ontologyParserQUIT, 0)
}

func (s *ExitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitExit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Exit() (localctx IExitContext) {
	this := p
	_ = this

	localctx = NewExitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ontologyParserRULE_exit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ontologyParserEXIT || _la == ontologyParserQUIT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Paradigm() IParadigmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParadigmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParadigmContext)
}

func (s *ExpressionContext) Process() IProcessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcessContext)
}

func (s *ExpressionContext) Local_configuration() ILocal_configurationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocal_configurationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocal_configurationContext)
}

func (s *ExpressionContext) Global_configuration() IGlobal_configurationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobal_configurationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobal_configurationContext)
}

func (s *ExpressionContext) Knowledge_base() IKnowledge_baseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKnowledge_baseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKnowledge_baseContext)
}

func (s *ExpressionContext) Load() ILoadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoadContext)
}

func (s *ExpressionContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ExpressionContext) Remove() IRemoveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemoveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemoveContext)
}

func (s *ExpressionContext) Show() IShowContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShowContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShowContext)
}

func (s *ExpressionContext) Mermaid() IMermaidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMermaidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMermaidContext)
}

func (s *ExpressionContext) Translate() ITranslateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITranslateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITranslateContext)
}

func (s *ExpressionContext) Traverse() ITraverseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITraverseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITraverseContext)
}

func (s *ExpressionContext) Configure() IConfigureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConfigureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConfigureContext)
}

func (s *ExpressionContext) Compose() IComposeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComposeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComposeContext)
}

func (s *ExpressionContext) Project() IProjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjectContext)
}

func (s *ExpressionContext) Info() IInfoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInfoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInfoContext)
}

func (s *ExpressionContext) Render() IRenderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRenderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRenderContext)
}

func (s *ExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(ontologyParserID, 0)
}

func (s *ExpressionContext) LPAR() antlr.TerminalNode {
	return s.GetToken(ontologyParserLPAR, 0)
}

func (s *ExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) RPAR() antlr.TerminalNode {
	return s.GetToken(ontologyParserRPAR, 0)
}

func (s *ExpressionContext) LSQ() antlr.TerminalNode {
	return s.GetToken(ontologyParserLSQ, 0)
}

func (s *ExpressionContext) INT() antlr.TerminalNode {
	return s.GetToken(ontologyParserINT, 0)
}

func (s *ExpressionContext) RSQ() antlr.TerminalNode {
	return s.GetToken(ontologyParserRSQ, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *ontologyParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, ontologyParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(104)
			p.Paradigm()
		}

	case 2:
		{
			p.SetState(105)
			p.Process()
		}

	case 3:
		{
			p.SetState(106)
			p.Local_configuration()
		}

	case 4:
		{
			p.SetState(107)
			p.Global_configuration()
		}

	case 5:
		{
			p.SetState(108)
			p.Knowledge_base()
		}

	case 6:
		{
			p.SetState(109)
			p.Load()
		}

	case 7:
		{
			p.SetState(110)
			p.Assignment()
		}

	case 8:
		{
			p.SetState(111)
			p.Remove()
		}

	case 9:
		{
			p.SetState(112)
			p.Show()
		}

	case 10:
		{
			p.SetState(113)
			p.Mermaid()
		}

	case 11:
		{
			p.SetState(114)
			p.Translate()
		}

	case 12:
		{
			p.SetState(115)
			p.Traverse()
		}

	case 13:
		{
			p.SetState(116)
			p.Configure()
		}

	case 14:
		{
			p.SetState(117)
			p.Compose()
		}

	case 15:
		{
			p.SetState(118)
			p.Project()
		}

	case 16:
		{
			p.SetState(119)
			p.Info()
		}

	case 17:
		{
			p.SetState(120)
			p.Render()
		}

	case 18:
		{
			p.SetState(121)
			p.Match(ontologyParserID)
		}

	case 19:
		{
			p.SetState(122)
			p.Match(ontologyParserLPAR)
		}
		{
			p.SetState(123)
			p.expression(0)
		}
		{
			p.SetState(124)
			p.Match(ontologyParserRPAR)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ontologyParserRULE_expression)
			p.SetState(128)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(129)
				p.Match(ontologyParserLSQ)
			}
			{
				p.SetState(130)
				p.Match(ontologyParserINT)
			}
			{
				p.SetState(131)
				p.Match(ontologyParserRSQ)
			}

		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// ILoadContext is an interface to support dynamic dispatch.
type ILoadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoadContext differentiates from other interfaces.
	IsLoadContext()
}

type LoadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoadContext() *LoadContext {
	var p = new(LoadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_load
	return p
}

func (*LoadContext) IsLoadContext() {}

func NewLoadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoadContext {
	var p = new(LoadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_load

	return p
}

func (s *LoadContext) GetParser() antlr.Parser { return s.parser }

func (s *LoadContext) LOAD() antlr.TerminalNode {
	return s.GetToken(ontologyParserLOAD, 0)
}

func (s *LoadContext) ID() antlr.TerminalNode {
	return s.GetToken(ontologyParserID, 0)
}

func (s *LoadContext) PATH() antlr.TerminalNode {
	return s.GetToken(ontologyParserPATH, 0)
}

func (s *LoadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitLoad(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Load() (localctx ILoadContext) {
	this := p
	_ = this

	localctx = NewLoadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ontologyParserRULE_load)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Match(ontologyParserLOAD)
	}
	{
		p.SetState(138)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ontologyParserPATH || _la == ontologyParserID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(ontologyParserID, 0)
}

func (s *AssignmentContext) ASGN() antlr.TerminalNode {
	return s.GetToken(ontologyParserASGN, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ontologyParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(141)
		p.Match(ontologyParserASGN)
	}
	{
		p.SetState(142)
		p.expression(0)
	}

	return localctx
}

// IRemoveContext is an interface to support dynamic dispatch.
type IRemoveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRemoveContext differentiates from other interfaces.
	IsRemoveContext()
}

type RemoveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRemoveContext() *RemoveContext {
	var p = new(RemoveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_remove
	return p
}

func (*RemoveContext) IsRemoveContext() {}

func NewRemoveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RemoveContext {
	var p = new(RemoveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_remove

	return p
}

func (s *RemoveContext) GetParser() antlr.Parser { return s.parser }

func (s *RemoveContext) REMOVE() antlr.TerminalNode {
	return s.GetToken(ontologyParserREMOVE, 0)
}

func (s *RemoveContext) ID() antlr.TerminalNode {
	return s.GetToken(ontologyParserID, 0)
}

func (s *RemoveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemoveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RemoveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitRemove(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Remove() (localctx IRemoveContext) {
	this := p
	_ = this

	localctx = NewRemoveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ontologyParserRULE_remove)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(ontologyParserREMOVE)
	}
	{
		p.SetState(145)
		p.Match(ontologyParserID)
	}

	return localctx
}

// IShowContext is an interface to support dynamic dispatch.
type IShowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShowContext differentiates from other interfaces.
	IsShowContext()
}

type ShowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShowContext() *ShowContext {
	var p = new(ShowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_show
	return p
}

func (*ShowContext) IsShowContext() {}

func NewShowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShowContext {
	var p = new(ShowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_show

	return p
}

func (s *ShowContext) GetParser() antlr.Parser { return s.parser }

func (s *ShowContext) SHOW() antlr.TerminalNode {
	return s.GetToken(ontologyParserSHOW, 0)
}

func (s *ShowContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ShowContext) LS() antlr.TerminalNode {
	return s.GetToken(ontologyParserLS, 0)
}

func (s *ShowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitShow(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Show() (localctx IShowContext) {
	this := p
	_ = this

	localctx = NewShowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ontologyParserRULE_show)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(155)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ontologyParserSHOW:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)
			p.Match(ontologyParserSHOW)
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(148)
				p.expression(0)
			}

		}

	case ontologyParserLS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(151)
			p.Match(ontologyParserLS)
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(152)
				p.expression(0)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMermaidContext is an interface to support dynamic dispatch.
type IMermaidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMermaidContext differentiates from other interfaces.
	IsMermaidContext()
}

type MermaidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMermaidContext() *MermaidContext {
	var p = new(MermaidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_mermaid
	return p
}

func (*MermaidContext) IsMermaidContext() {}

func NewMermaidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MermaidContext {
	var p = new(MermaidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_mermaid

	return p
}

func (s *MermaidContext) GetParser() antlr.Parser { return s.parser }

func (s *MermaidContext) MERMAID() antlr.TerminalNode {
	return s.GetToken(ontologyParserMERMAID, 0)
}

func (s *MermaidContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MermaidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MermaidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MermaidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitMermaid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Mermaid() (localctx IMermaidContext) {
	this := p
	_ = this

	localctx = NewMermaidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ontologyParserRULE_mermaid)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(ontologyParserMERMAID)
	}
	{
		p.SetState(158)
		p.expression(0)
	}

	return localctx
}

// ITranslateContext is an interface to support dynamic dispatch.
type ITranslateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTranslateContext differentiates from other interfaces.
	IsTranslateContext()
}

type TranslateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTranslateContext() *TranslateContext {
	var p = new(TranslateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_translate
	return p
}

func (*TranslateContext) IsTranslateContext() {}

func NewTranslateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TranslateContext {
	var p = new(TranslateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_translate

	return p
}

func (s *TranslateContext) GetParser() antlr.Parser { return s.parser }

func (s *TranslateContext) TRANSLATE() antlr.TerminalNode {
	return s.GetToken(ontologyParserTRANSLATE, 0)
}

func (s *TranslateContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TranslateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TranslateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TranslateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitTranslate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Translate() (localctx ITranslateContext) {
	this := p
	_ = this

	localctx = NewTranslateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ontologyParserRULE_translate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(ontologyParserTRANSLATE)
	}
	{
		p.SetState(161)
		p.expression(0)
	}

	return localctx
}

// ITraverseContext is an interface to support dynamic dispatch.
type ITraverseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTraverseContext differentiates from other interfaces.
	IsTraverseContext()
}

type TraverseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTraverseContext() *TraverseContext {
	var p = new(TraverseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_traverse
	return p
}

func (*TraverseContext) IsTraverseContext() {}

func NewTraverseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TraverseContext {
	var p = new(TraverseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_traverse

	return p
}

func (s *TraverseContext) GetParser() antlr.Parser { return s.parser }

func (s *TraverseContext) TRAVERSE() antlr.TerminalNode {
	return s.GetToken(ontologyParserTRAVERSE, 0)
}

func (s *TraverseContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *TraverseContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *TraverseContext) DOT() antlr.TerminalNode {
	return s.GetToken(ontologyParserDOT, 0)
}

func (s *TraverseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TraverseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TraverseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TraverseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitTraverse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Traverse() (localctx ITraverseContext) {
	this := p
	_ = this

	localctx = NewTraverseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ontologyParserRULE_traverse)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(ontologyParserTRAVERSE)
	}
	{
		p.SetState(164)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(165)
		p.Match(ontologyParserDOT)
	}
	{
		p.SetState(166)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(167)
		p.expression(0)
	}

	return localctx
}

// IConfigureContext is an interface to support dynamic dispatch.
type IConfigureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConfigureContext differentiates from other interfaces.
	IsConfigureContext()
}

type ConfigureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigureContext() *ConfigureContext {
	var p = new(ConfigureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_configure
	return p
}

func (*ConfigureContext) IsConfigureContext() {}

func NewConfigureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigureContext {
	var p = new(ConfigureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_configure

	return p
}

func (s *ConfigureContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigureContext) CONFIGURE() antlr.TerminalNode {
	return s.GetToken(ontologyParserCONFIGURE, 0)
}

func (s *ConfigureContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ConfigureContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConfigureContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *ConfigureContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *ConfigureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitConfigure(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Configure() (localctx IConfigureContext) {
	this := p
	_ = this

	localctx = NewConfigureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ontologyParserRULE_configure)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(ontologyParserCONFIGURE)
	}
	{
		p.SetState(170)
		p.expression(0)
	}
	{
		p.SetState(171)
		p.expression(0)
	}
	{
		p.SetState(172)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(173)
		p.Match(ontologyParserID)
	}

	return localctx
}

// IComposeContext is an interface to support dynamic dispatch.
type IComposeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComposeContext differentiates from other interfaces.
	IsComposeContext()
}

type ComposeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComposeContext() *ComposeContext {
	var p = new(ComposeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_compose
	return p
}

func (*ComposeContext) IsComposeContext() {}

func NewComposeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComposeContext {
	var p = new(ComposeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_compose

	return p
}

func (s *ComposeContext) GetParser() antlr.Parser { return s.parser }

func (s *ComposeContext) COMPOSE() antlr.TerminalNode {
	return s.GetToken(ontologyParserCOMPOSE, 0)
}

func (s *ComposeContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ComposeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComposeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComposeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitCompose(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Compose() (localctx IComposeContext) {
	this := p
	_ = this

	localctx = NewComposeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ontologyParserRULE_compose)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(ontologyParserCOMPOSE)
	}
	{
		p.SetState(176)
		p.expression(0)
	}

	return localctx
}

// IProjectContext is an interface to support dynamic dispatch.
type IProjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProjectContext differentiates from other interfaces.
	IsProjectContext()
}

type ProjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectContext() *ProjectContext {
	var p = new(ProjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_project
	return p
}

func (*ProjectContext) IsProjectContext() {}

func NewProjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectContext {
	var p = new(ProjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_project

	return p
}

func (s *ProjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectContext) PROJECT() antlr.TerminalNode {
	return s.GetToken(ontologyParserPROJECT, 0)
}

func (s *ProjectContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitProject(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Project() (localctx IProjectContext) {
	this := p
	_ = this

	localctx = NewProjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ontologyParserRULE_project)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(ontologyParserPROJECT)
	}
	{
		p.SetState(179)
		p.expression(0)
	}

	return localctx
}

// IInfoContext is an interface to support dynamic dispatch.
type IInfoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInfoContext differentiates from other interfaces.
	IsInfoContext()
}

type InfoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInfoContext() *InfoContext {
	var p = new(InfoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_info
	return p
}

func (*InfoContext) IsInfoContext() {}

func NewInfoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InfoContext {
	var p = new(InfoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_info

	return p
}

func (s *InfoContext) GetParser() antlr.Parser { return s.parser }

func (s *InfoContext) INFO() antlr.TerminalNode {
	return s.GetToken(ontologyParserINFO, 0)
}

func (s *InfoContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InfoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InfoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitInfo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Info() (localctx IInfoContext) {
	this := p
	_ = this

	localctx = NewInfoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ontologyParserRULE_info)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(ontologyParserINFO)
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(182)
			p.expression(0)
		}

	}

	return localctx
}

// IRenderContext is an interface to support dynamic dispatch.
type IRenderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRenderContext differentiates from other interfaces.
	IsRenderContext()
}

type RenderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRenderContext() *RenderContext {
	var p = new(RenderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_render
	return p
}

func (*RenderContext) IsRenderContext() {}

func NewRenderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RenderContext {
	var p = new(RenderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_render

	return p
}

func (s *RenderContext) GetParser() antlr.Parser { return s.parser }

func (s *RenderContext) RENDER() antlr.TerminalNode {
	return s.GetToken(ontologyParserRENDER, 0)
}

func (s *RenderContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RenderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RenderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RenderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitRender(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Render() (localctx IRenderContext) {
	this := p
	_ = this

	localctx = NewRenderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ontologyParserRULE_render)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(ontologyParserRENDER)
	}
	{
		p.SetState(186)
		p.expression(0)
	}

	return localctx
}

// IParadigmContext is an interface to support dynamic dispatch.
type IParadigmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParadigmContext differentiates from other interfaces.
	IsParadigmContext()
}

type ParadigmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParadigmContext() *ParadigmContext {
	var p = new(ParadigmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_paradigm
	return p
}

func (*ParadigmContext) IsParadigmContext() {}

func NewParadigmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParadigmContext {
	var p = new(ParadigmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_paradigm

	return p
}

func (s *ParadigmContext) GetParser() antlr.Parser { return s.parser }

func (s *ParadigmContext) PARADIGM() antlr.TerminalNode {
	return s.GetToken(ontologyParserPARADIGM, 0)
}

func (s *ParadigmContext) LBRA() antlr.TerminalNode {
	return s.GetToken(ontologyParserLBRA, 0)
}

func (s *ParadigmContext) RBRA() antlr.TerminalNode {
	return s.GetToken(ontologyParserRBRA, 0)
}

func (s *ParadigmContext) AllParadigm_declaration() []IParadigm_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParadigm_declarationContext); ok {
			len++
		}
	}

	tst := make([]IParadigm_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParadigm_declarationContext); ok {
			tst[i] = t.(IParadigm_declarationContext)
			i++
		}
	}

	return tst
}

func (s *ParadigmContext) Paradigm_declaration(i int) IParadigm_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParadigm_declarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParadigm_declarationContext)
}

func (s *ParadigmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParadigmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParadigmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitParadigm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Paradigm() (localctx IParadigmContext) {
	this := p
	_ = this

	localctx = NewParadigmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ontologyParserRULE_paradigm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(ontologyParserPARADIGM)
	}
	{
		p.SetState(189)
		p.Match(ontologyParserLBRA)
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ontologyParserPROPERTY)|(1<<ontologyParserMODEL)|(1<<ontologyParserPHYSICAL)|(1<<ontologyParserACTUATOR)|(1<<ontologyParserTRANSLATION))) != 0 {
		{
			p.SetState(190)
			p.Paradigm_declaration()
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(196)
		p.Match(ontologyParserRBRA)
	}

	return localctx
}

// IParadigm_declarationContext is an interface to support dynamic dispatch.
type IParadigm_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParadigm_declarationContext differentiates from other interfaces.
	IsParadigm_declarationContext()
}

type Paradigm_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParadigm_declarationContext() *Paradigm_declarationContext {
	var p = new(Paradigm_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_paradigm_declaration
	return p
}

func (*Paradigm_declarationContext) IsParadigm_declarationContext() {}

func NewParadigm_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Paradigm_declarationContext {
	var p = new(Paradigm_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_paradigm_declaration

	return p
}

func (s *Paradigm_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Paradigm_declarationContext) Property() IPropertyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *Paradigm_declarationContext) Class() IClassContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassContext)
}

func (s *Paradigm_declarationContext) Model() IModelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModelContext)
}

func (s *Paradigm_declarationContext) Translation() ITranslationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITranslationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITranslationContext)
}

func (s *Paradigm_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Paradigm_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Paradigm_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitParadigm_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Paradigm_declaration() (localctx IParadigm_declarationContext) {
	this := p
	_ = this

	localctx = NewParadigm_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ontologyParserRULE_paradigm_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(202)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ontologyParserPROPERTY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(198)
			p.Property()
		}

	case ontologyParserPHYSICAL, ontologyParserACTUATOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(199)
			p.Class()
		}

	case ontologyParserMODEL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(200)
			p.Model()
		}

	case ontologyParserTRANSLATION:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(201)
			p.Translation()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_property
	return p
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) PROPERTY() antlr.TerminalNode {
	return s.GetToken(ontologyParserPROPERTY, 0)
}

func (s *PropertyContext) AllType() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *PropertyContext) Type(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *PropertyContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserCOMMA)
}

func (s *PropertyContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserCOMMA, i)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Property() (localctx IPropertyContext) {
	this := p
	_ = this

	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ontologyParserRULE_property)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(ontologyParserPROPERTY)
	}
	{
		p.SetState(205)
		p.Type()
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ontologyParserCOMMA {
		{
			p.SetState(206)
			p.Match(ontologyParserCOMMA)
		}
		{
			p.SetState(207)
			p.Type()
		}

		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *TypeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *TypeContext) LBRA() antlr.TerminalNode {
	return s.GetToken(ontologyParserLBRA, 0)
}

func (s *TypeContext) RBRA() antlr.TerminalNode {
	return s.GetToken(ontologyParserRBRA, 0)
}

func (s *TypeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserCOMMA)
}

func (s *TypeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserCOMMA, i)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Type() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ontologyParserRULE_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(213)
			p.Match(ontologyParserID)
		}
		{
			p.SetState(214)
			p.Match(ontologyParserLBRA)
		}
		{
			p.SetState(215)
			p.Match(ontologyParserID)
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ontologyParserCOMMA {
			{
				p.SetState(216)
				p.Match(ontologyParserCOMMA)
			}
			{
				p.SetState(217)
				p.Match(ontologyParserID)
			}

			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(223)
			p.Match(ontologyParserRBRA)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)
			p.Match(ontologyParserID)
		}

	}

	return localctx
}

// IModelContext is an interface to support dynamic dispatch.
type IModelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModelContext differentiates from other interfaces.
	IsModelContext()
}

type ModelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelContext() *ModelContext {
	var p = new(ModelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_model
	return p
}

func (*ModelContext) IsModelContext() {}

func NewModelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelContext {
	var p = new(ModelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_model

	return p
}

func (s *ModelContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelContext) MODEL() antlr.TerminalNode {
	return s.GetToken(ontologyParserMODEL, 0)
}

func (s *ModelContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *ModelContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *ModelContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserCOMMA)
}

func (s *ModelContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserCOMMA, i)
}

func (s *ModelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitModel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Model() (localctx IModelContext) {
	this := p
	_ = this

	localctx = NewModelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ontologyParserRULE_model)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(ontologyParserMODEL)
	}
	{
		p.SetState(228)
		p.Match(ontologyParserID)
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ontologyParserCOMMA {
		{
			p.SetState(229)
			p.Match(ontologyParserCOMMA)
		}
		{
			p.SetState(230)
			p.Match(ontologyParserID)
		}

		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IClassContext is an interface to support dynamic dispatch.
type IClassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassContext differentiates from other interfaces.
	IsClassContext()
}

type ClassContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassContext() *ClassContext {
	var p = new(ClassContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_class
	return p
}

func (*ClassContext) IsClassContext() {}

func NewClassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassContext {
	var p = new(ClassContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_class

	return p
}

func (s *ClassContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *ClassContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *ClassContext) LPAR() antlr.TerminalNode {
	return s.GetToken(ontologyParserLPAR, 0)
}

func (s *ClassContext) RPAR() antlr.TerminalNode {
	return s.GetToken(ontologyParserRPAR, 0)
}

func (s *ClassContext) COLON() antlr.TerminalNode {
	return s.GetToken(ontologyParserCOLON, 0)
}

func (s *ClassContext) AllInternal_edge() []IInternal_edgeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInternal_edgeContext); ok {
			len++
		}
	}

	tst := make([]IInternal_edgeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInternal_edgeContext); ok {
			tst[i] = t.(IInternal_edgeContext)
			i++
		}
	}

	return tst
}

func (s *ClassContext) Internal_edge(i int) IInternal_edgeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInternal_edgeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInternal_edgeContext)
}

func (s *ClassContext) PHYSICAL() antlr.TerminalNode {
	return s.GetToken(ontologyParserPHYSICAL, 0)
}

func (s *ClassContext) ACTUATOR() antlr.TerminalNode {
	return s.GetToken(ontologyParserACTUATOR, 0)
}

func (s *ClassContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserCOMMA)
}

func (s *ClassContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserCOMMA, i)
}

func (s *ClassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitClass(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Class() (localctx IClassContext) {
	this := p
	_ = this

	localctx = NewClassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ontologyParserRULE_class)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ontologyParserPHYSICAL || _la == ontologyParserACTUATOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(237)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(238)
		p.Match(ontologyParserLPAR)
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ontologyParserID {
		{
			p.SetState(239)
			p.Match(ontologyParserID)
		}
		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ontologyParserCOMMA {
			{
				p.SetState(240)
				p.Match(ontologyParserCOMMA)
			}
			{
				p.SetState(241)
				p.Match(ontologyParserID)
			}

			p.SetState(246)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(249)
		p.Match(ontologyParserRPAR)
	}
	{
		p.SetState(250)
		p.Match(ontologyParserCOLON)
	}
	{
		p.SetState(251)
		p.Internal_edge()
	}

	{
		p.SetState(252)
		p.Match(ontologyParserCOMMA)
	}
	{
		p.SetState(253)
		p.Internal_edge()
	}

	return localctx
}

// IInternal_edgeContext is an interface to support dynamic dispatch.
type IInternal_edgeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInternal_edgeContext differentiates from other interfaces.
	IsInternal_edgeContext()
}

type Internal_edgeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInternal_edgeContext() *Internal_edgeContext {
	var p = new(Internal_edgeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_internal_edge
	return p
}

func (*Internal_edgeContext) IsInternal_edgeContext() {}

func NewInternal_edgeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Internal_edgeContext {
	var p = new(Internal_edgeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_internal_edge

	return p
}

func (s *Internal_edgeContext) GetParser() antlr.Parser { return s.parser }

func (s *Internal_edgeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *Internal_edgeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *Internal_edgeContext) ARROW() antlr.TerminalNode {
	return s.GetToken(ontologyParserARROW, 0)
}

func (s *Internal_edgeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Internal_edgeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Internal_edgeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitInternal_edge(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Internal_edge() (localctx IInternal_edgeContext) {
	this := p
	_ = this

	localctx = NewInternal_edgeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ontologyParserRULE_internal_edge)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(256)
		p.Match(ontologyParserARROW)
	}
	{
		p.SetState(257)
		p.Match(ontologyParserID)
	}

	return localctx
}

// ITranslationContext is an interface to support dynamic dispatch.
type ITranslationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTranslationContext differentiates from other interfaces.
	IsTranslationContext()
}

type TranslationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTranslationContext() *TranslationContext {
	var p = new(TranslationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_translation
	return p
}

func (*TranslationContext) IsTranslationContext() {}

func NewTranslationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TranslationContext {
	var p = new(TranslationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_translation

	return p
}

func (s *TranslationContext) GetParser() antlr.Parser { return s.parser }

func (s *TranslationContext) TRANSLATION() antlr.TerminalNode {
	return s.GetToken(ontologyParserTRANSLATION, 0)
}

func (s *TranslationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *TranslationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *TranslationContext) ARROW() antlr.TerminalNode {
	return s.GetToken(ontologyParserARROW, 0)
}

func (s *TranslationContext) COLON() antlr.TerminalNode {
	return s.GetToken(ontologyParserCOLON, 0)
}

func (s *TranslationContext) AllArg_connection() []IArg_connectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArg_connectionContext); ok {
			len++
		}
	}

	tst := make([]IArg_connectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArg_connectionContext); ok {
			tst[i] = t.(IArg_connectionContext)
			i++
		}
	}

	return tst
}

func (s *TranslationContext) Arg_connection(i int) IArg_connectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArg_connectionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArg_connectionContext)
}

func (s *TranslationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserCOMMA)
}

func (s *TranslationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserCOMMA, i)
}

func (s *TranslationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TranslationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TranslationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitTranslation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Translation() (localctx ITranslationContext) {
	this := p
	_ = this

	localctx = NewTranslationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ontologyParserRULE_translation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(ontologyParserTRANSLATION)
	}
	{
		p.SetState(260)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(261)
		p.Match(ontologyParserARROW)
	}
	{
		p.SetState(262)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(263)
		p.Match(ontologyParserCOLON)
	}
	{
		p.SetState(264)
		p.Arg_connection()
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ontologyParserCOMMA {
		{
			p.SetState(265)
			p.Match(ontologyParserCOMMA)
		}
		{
			p.SetState(266)
			p.Arg_connection()
		}

		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArg_connectionContext is an interface to support dynamic dispatch.
type IArg_connectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArg_connectionContext differentiates from other interfaces.
	IsArg_connectionContext()
}

type Arg_connectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArg_connectionContext() *Arg_connectionContext {
	var p = new(Arg_connectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_arg_connection
	return p
}

func (*Arg_connectionContext) IsArg_connectionContext() {}

func NewArg_connectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arg_connectionContext {
	var p = new(Arg_connectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_arg_connection

	return p
}

func (s *Arg_connectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Arg_connectionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *Arg_connectionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *Arg_connectionContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserDOT)
}

func (s *Arg_connectionContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserDOT, i)
}

func (s *Arg_connectionContext) ARROW() antlr.TerminalNode {
	return s.GetToken(ontologyParserARROW, 0)
}

func (s *Arg_connectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_connectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arg_connectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitArg_connection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Arg_connection() (localctx IArg_connectionContext) {
	this := p
	_ = this

	localctx = NewArg_connectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ontologyParserRULE_arg_connection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(273)
		p.Match(ontologyParserDOT)
	}
	{
		p.SetState(274)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(275)
		p.Match(ontologyParserARROW)
	}
	{
		p.SetState(276)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(277)
		p.Match(ontologyParserDOT)
	}
	{
		p.SetState(278)
		p.Match(ontologyParserID)
	}

	return localctx
}

// IConnectionContext is an interface to support dynamic dispatch.
type IConnectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConnectionContext differentiates from other interfaces.
	IsConnectionContext()
}

type ConnectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConnectionContext() *ConnectionContext {
	var p = new(ConnectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_connection
	return p
}

func (*ConnectionContext) IsConnectionContext() {}

func NewConnectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConnectionContext {
	var p = new(ConnectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_connection

	return p
}

func (s *ConnectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConnectionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *ConnectionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *ConnectionContext) ARROW() antlr.TerminalNode {
	return s.GetToken(ontologyParserARROW, 0)
}

func (s *ConnectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConnectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitConnection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Connection() (localctx IConnectionContext) {
	this := p
	_ = this

	localctx = NewConnectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ontologyParserRULE_connection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(281)
		p.Match(ontologyParserARROW)
	}
	{
		p.SetState(282)
		p.Match(ontologyParserID)
	}

	return localctx
}

// IProcessContext is an interface to support dynamic dispatch.
type IProcessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcessContext differentiates from other interfaces.
	IsProcessContext()
}

type ProcessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcessContext() *ProcessContext {
	var p = new(ProcessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_process
	return p
}

func (*ProcessContext) IsProcessContext() {}

func NewProcessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcessContext {
	var p = new(ProcessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_process

	return p
}

func (s *ProcessContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcessContext) PROCESS() antlr.TerminalNode {
	return s.GetToken(ontologyParserPROCESS, 0)
}

func (s *ProcessContext) ID() antlr.TerminalNode {
	return s.GetToken(ontologyParserID, 0)
}

func (s *ProcessContext) LBRA() antlr.TerminalNode {
	return s.GetToken(ontologyParserLBRA, 0)
}

func (s *ProcessContext) RBRA() antlr.TerminalNode {
	return s.GetToken(ontologyParserRBRA, 0)
}

func (s *ProcessContext) AllProcess_declaration() []IProcess_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProcess_declarationContext); ok {
			len++
		}
	}

	tst := make([]IProcess_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProcess_declarationContext); ok {
			tst[i] = t.(IProcess_declarationContext)
			i++
		}
	}

	return tst
}

func (s *ProcessContext) Process_declaration(i int) IProcess_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcess_declarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcess_declarationContext)
}

func (s *ProcessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitProcess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Process() (localctx IProcessContext) {
	this := p
	_ = this

	localctx = NewProcessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ontologyParserRULE_process)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Match(ontologyParserPROCESS)
	}
	{
		p.SetState(285)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(286)
		p.Match(ontologyParserLBRA)
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ontologyParserPHYSICAL)|(1<<ontologyParserACTUATOR)|(1<<ontologyParserSENSOR)|(1<<ontologyParserDEVICE)|(1<<ontologyParserCONNECTION))) != 0 {
		{
			p.SetState(287)
			p.Process_declaration()
		}

		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(293)
		p.Match(ontologyParserRBRA)
	}

	return localctx
}

// IProcess_declarationContext is an interface to support dynamic dispatch.
type IProcess_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcess_declarationContext differentiates from other interfaces.
	IsProcess_declarationContext()
}

type Process_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcess_declarationContext() *Process_declarationContext {
	var p = new(Process_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_process_declaration
	return p
}

func (*Process_declarationContext) IsProcess_declarationContext() {}

func NewProcess_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Process_declarationContext {
	var p = new(Process_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_process_declaration

	return p
}

func (s *Process_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Process_declarationContext) Device() IDeviceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeviceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeviceContext)
}

func (s *Process_declarationContext) Component() IComponentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComponentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComponentContext)
}

func (s *Process_declarationContext) Connection_decl() IConnection_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConnection_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConnection_declContext)
}

func (s *Process_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Process_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Process_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitProcess_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Process_declaration() (localctx IProcess_declarationContext) {
	this := p
	_ = this

	localctx = NewProcess_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ontologyParserRULE_process_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(298)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ontologyParserDEVICE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(295)
			p.Device()
		}

	case ontologyParserPHYSICAL, ontologyParserACTUATOR, ontologyParserSENSOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(296)
			p.Component()
		}

	case ontologyParserCONNECTION:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(297)
			p.Connection_decl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDeviceContext is an interface to support dynamic dispatch.
type IDeviceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeviceContext differentiates from other interfaces.
	IsDeviceContext()
}

type DeviceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeviceContext() *DeviceContext {
	var p = new(DeviceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_device
	return p
}

func (*DeviceContext) IsDeviceContext() {}

func NewDeviceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeviceContext {
	var p = new(DeviceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_device

	return p
}

func (s *DeviceContext) GetParser() antlr.Parser { return s.parser }

func (s *DeviceContext) DEVICE() antlr.TerminalNode {
	return s.GetToken(ontologyParserDEVICE, 0)
}

func (s *DeviceContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *DeviceContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *DeviceContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserCOMMA)
}

func (s *DeviceContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserCOMMA, i)
}

func (s *DeviceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeviceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeviceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitDevice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Device() (localctx IDeviceContext) {
	this := p
	_ = this

	localctx = NewDeviceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ontologyParserRULE_device)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(ontologyParserDEVICE)
	}
	{
		p.SetState(301)
		p.Match(ontologyParserID)
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ontologyParserCOMMA {
		{
			p.SetState(302)
			p.Match(ontologyParserCOMMA)
		}
		{
			p.SetState(303)
			p.Match(ontologyParserID)
		}

		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IComponentContext is an interface to support dynamic dispatch.
type IComponentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentContext differentiates from other interfaces.
	IsComponentContext()
}

type ComponentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentContext() *ComponentContext {
	var p = new(ComponentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_component
	return p
}

func (*ComponentContext) IsComponentContext() {}

func NewComponentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentContext {
	var p = new(ComponentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_component

	return p
}

func (s *ComponentContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentContext) PHYSICAL() antlr.TerminalNode {
	return s.GetToken(ontologyParserPHYSICAL, 0)
}

func (s *ComponentContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *ComponentContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *ComponentContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserCOMMA)
}

func (s *ComponentContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserCOMMA, i)
}

func (s *ComponentContext) ACTUATOR() antlr.TerminalNode {
	return s.GetToken(ontologyParserACTUATOR, 0)
}

func (s *ComponentContext) AllAT() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserAT)
}

func (s *ComponentContext) AT(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserAT, i)
}

func (s *ComponentContext) SENSOR() antlr.TerminalNode {
	return s.GetToken(ontologyParserSENSOR, 0)
}

func (s *ComponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitComponent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Component() (localctx IComponentContext) {
	this := p
	_ = this

	localctx = NewComponentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ontologyParserRULE_component)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(347)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ontologyParserPHYSICAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(309)
			p.Match(ontologyParserPHYSICAL)
		}
		{
			p.SetState(310)
			p.Match(ontologyParserID)
		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ontologyParserCOMMA {
			{
				p.SetState(311)
				p.Match(ontologyParserCOMMA)
			}
			{
				p.SetState(312)
				p.Match(ontologyParserID)
			}

			p.SetState(317)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(318)
			p.Match(ontologyParserID)
		}

	case ontologyParserACTUATOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(319)
			p.Match(ontologyParserACTUATOR)
		}
		{
			p.SetState(320)
			p.Match(ontologyParserID)
		}
		{
			p.SetState(321)
			p.Match(ontologyParserAT)
		}
		{
			p.SetState(322)
			p.Match(ontologyParserID)
		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ontologyParserCOMMA {
			{
				p.SetState(323)
				p.Match(ontologyParserCOMMA)
			}
			{
				p.SetState(324)
				p.Match(ontologyParserID)
			}
			{
				p.SetState(325)
				p.Match(ontologyParserAT)
			}
			{
				p.SetState(326)
				p.Match(ontologyParserID)
			}

			p.SetState(331)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(332)
			p.Match(ontologyParserID)
		}

	case ontologyParserSENSOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(333)
			p.Match(ontologyParserSENSOR)
		}
		{
			p.SetState(334)
			p.Match(ontologyParserID)
		}
		{
			p.SetState(335)
			p.Match(ontologyParserAT)
		}
		{
			p.SetState(336)
			p.Match(ontologyParserID)
		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ontologyParserCOMMA {
			{
				p.SetState(337)
				p.Match(ontologyParserCOMMA)
			}
			{
				p.SetState(338)
				p.Match(ontologyParserID)
			}
			{
				p.SetState(339)
				p.Match(ontologyParserAT)
			}
			{
				p.SetState(340)
				p.Match(ontologyParserID)
			}

			p.SetState(345)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(346)
			p.Match(ontologyParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConnection_declContext is an interface to support dynamic dispatch.
type IConnection_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConnection_declContext differentiates from other interfaces.
	IsConnection_declContext()
}

type Connection_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConnection_declContext() *Connection_declContext {
	var p = new(Connection_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_connection_decl
	return p
}

func (*Connection_declContext) IsConnection_declContext() {}

func NewConnection_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Connection_declContext {
	var p = new(Connection_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_connection_decl

	return p
}

func (s *Connection_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Connection_declContext) CONNECTION() antlr.TerminalNode {
	return s.GetToken(ontologyParserCONNECTION, 0)
}

func (s *Connection_declContext) AllConnection() []IConnectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConnectionContext); ok {
			len++
		}
	}

	tst := make([]IConnectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConnectionContext); ok {
			tst[i] = t.(IConnectionContext)
			i++
		}
	}

	return tst
}

func (s *Connection_declContext) Connection(i int) IConnectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConnectionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConnectionContext)
}

func (s *Connection_declContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserCOMMA)
}

func (s *Connection_declContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserCOMMA, i)
}

func (s *Connection_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Connection_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Connection_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitConnection_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Connection_decl() (localctx IConnection_declContext) {
	this := p
	_ = this

	localctx = NewConnection_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ontologyParserRULE_connection_decl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.Match(ontologyParserCONNECTION)
	}
	{
		p.SetState(350)
		p.Connection()
	}
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ontologyParserCOMMA {
		{
			p.SetState(351)
			p.Match(ontologyParserCOMMA)
		}
		{
			p.SetState(352)
			p.Connection()
		}

		p.SetState(357)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILocal_configurationContext is an interface to support dynamic dispatch.
type ILocal_configurationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocal_configurationContext differentiates from other interfaces.
	IsLocal_configurationContext()
}

type Local_configurationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocal_configurationContext() *Local_configurationContext {
	var p = new(Local_configurationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_local_configuration
	return p
}

func (*Local_configurationContext) IsLocal_configurationContext() {}

func NewLocal_configurationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Local_configurationContext {
	var p = new(Local_configurationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_local_configuration

	return p
}

func (s *Local_configurationContext) GetParser() antlr.Parser { return s.parser }

func (s *Local_configurationContext) LOCAL() antlr.TerminalNode {
	return s.GetToken(ontologyParserLOCAL, 0)
}

func (s *Local_configurationContext) LBRA() antlr.TerminalNode {
	return s.GetToken(ontologyParserLBRA, 0)
}

func (s *Local_configurationContext) RBRA() antlr.TerminalNode {
	return s.GetToken(ontologyParserRBRA, 0)
}

func (s *Local_configurationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *Local_configurationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *Local_configurationContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserEQ)
}

func (s *Local_configurationContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserEQ, i)
}

func (s *Local_configurationContext) AllLocal() []ILocalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILocalContext); ok {
			len++
		}
	}

	tst := make([]ILocalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILocalContext); ok {
			tst[i] = t.(ILocalContext)
			i++
		}
	}

	return tst
}

func (s *Local_configurationContext) Local(i int) ILocalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalContext)
}

func (s *Local_configurationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Local_configurationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Local_configurationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitLocal_configuration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Local_configuration() (localctx ILocal_configurationContext) {
	this := p
	_ = this

	localctx = NewLocal_configurationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ontologyParserRULE_local_configuration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(358)
		p.Match(ontologyParserLOCAL)
	}
	{
		p.SetState(359)
		p.Match(ontologyParserLBRA)
	}
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ontologyParserID {
		{
			p.SetState(360)
			p.Match(ontologyParserID)
		}
		{
			p.SetState(361)
			p.Match(ontologyParserEQ)
		}
		{
			p.SetState(362)
			p.Local()
		}

		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(367)
		p.Match(ontologyParserRBRA)
	}

	return localctx
}

// IGlobal_configurationContext is an interface to support dynamic dispatch.
type IGlobal_configurationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGlobal_configurationContext differentiates from other interfaces.
	IsGlobal_configurationContext()
}

type Global_configurationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobal_configurationContext() *Global_configurationContext {
	var p = new(Global_configurationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_global_configuration
	return p
}

func (*Global_configurationContext) IsGlobal_configurationContext() {}

func NewGlobal_configurationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Global_configurationContext {
	var p = new(Global_configurationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_global_configuration

	return p
}

func (s *Global_configurationContext) GetParser() antlr.Parser { return s.parser }

func (s *Global_configurationContext) GLOBAL() antlr.TerminalNode {
	return s.GetToken(ontologyParserGLOBAL, 0)
}

func (s *Global_configurationContext) Global() IGlobalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobalContext)
}

func (s *Global_configurationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Global_configurationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Global_configurationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitGlobal_configuration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Global_configuration() (localctx IGlobal_configurationContext) {
	this := p
	_ = this

	localctx = NewGlobal_configurationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ontologyParserRULE_global_configuration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		p.Match(ontologyParserGLOBAL)
	}
	{
		p.SetState(370)
		p.Global()
	}

	return localctx
}

// ILocalContext is an interface to support dynamic dispatch.
type ILocalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalContext differentiates from other interfaces.
	IsLocalContext()
}

type LocalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalContext() *LocalContext {
	var p = new(LocalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_local
	return p
}

func (*LocalContext) IsLocalContext() {}

func NewLocalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalContext {
	var p = new(LocalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_local

	return p
}

func (s *LocalContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalContext) END() antlr.TerminalNode {
	return s.GetToken(ontologyParserEND, 0)
}

func (s *LocalContext) Send() ISendContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISendContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISendContext)
}

func (s *LocalContext) Receive() IReceiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReceiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReceiveContext)
}

func (s *LocalContext) Loop() ILoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *LocalContext) Label() ILabelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LocalContext) Select() ISelectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectContext)
}

func (s *LocalContext) Branch() IBranchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBranchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBranchContext)
}

func (s *LocalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitLocal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Local() (localctx ILocalContext) {
	this := p
	_ = this

	localctx = NewLocalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ontologyParserRULE_local)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(372)
			p.Match(ontologyParserEND)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(373)
			p.Send()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(374)
			p.Receive()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(375)
			p.Loop()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(376)
			p.Label()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(377)
			p.Select()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(378)
			p.Branch()
		}

	}

	return localctx
}

// ISendContext is an interface to support dynamic dispatch.
type ISendContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSendContext differentiates from other interfaces.
	IsSendContext()
}

type SendContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySendContext() *SendContext {
	var p = new(SendContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_send
	return p
}

func (*SendContext) IsSendContext() {}

func NewSendContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SendContext {
	var p = new(SendContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_send

	return p
}

func (s *SendContext) GetParser() antlr.Parser { return s.parser }

func (s *SendContext) ID() antlr.TerminalNode {
	return s.GetToken(ontologyParserID, 0)
}

func (s *SendContext) BANG() antlr.TerminalNode {
	return s.GetToken(ontologyParserBANG, 0)
}

func (s *SendContext) Type() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *SendContext) DOT() antlr.TerminalNode {
	return s.GetToken(ontologyParserDOT, 0)
}

func (s *SendContext) Local() ILocalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalContext)
}

func (s *SendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitSend(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Send() (localctx ISendContext) {
	this := p
	_ = this

	localctx = NewSendContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ontologyParserRULE_send)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(382)
		p.Match(ontologyParserBANG)
	}
	{
		p.SetState(383)
		p.Type()
	}
	{
		p.SetState(384)
		p.Match(ontologyParserDOT)
	}
	{
		p.SetState(385)
		p.Local()
	}

	return localctx
}

// IReceiveContext is an interface to support dynamic dispatch.
type IReceiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReceiveContext differentiates from other interfaces.
	IsReceiveContext()
}

type ReceiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReceiveContext() *ReceiveContext {
	var p = new(ReceiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_receive
	return p
}

func (*ReceiveContext) IsReceiveContext() {}

func NewReceiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReceiveContext {
	var p = new(ReceiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_receive

	return p
}

func (s *ReceiveContext) GetParser() antlr.Parser { return s.parser }

func (s *ReceiveContext) ID() antlr.TerminalNode {
	return s.GetToken(ontologyParserID, 0)
}

func (s *ReceiveContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(ontologyParserQUESTION, 0)
}

func (s *ReceiveContext) Type() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ReceiveContext) DOT() antlr.TerminalNode {
	return s.GetToken(ontologyParserDOT, 0)
}

func (s *ReceiveContext) Local() ILocalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalContext)
}

func (s *ReceiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReceiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReceiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitReceive(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Receive() (localctx IReceiveContext) {
	this := p
	_ = this

	localctx = NewReceiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ontologyParserRULE_receive)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(388)
		p.Match(ontologyParserQUESTION)
	}
	{
		p.SetState(389)
		p.Type()
	}
	{
		p.SetState(390)
		p.Match(ontologyParserDOT)
	}
	{
		p.SetState(391)
		p.Local()
	}

	return localctx
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_loop
	return p
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) ID() antlr.TerminalNode {
	return s.GetToken(ontologyParserID, 0)
}

func (s *LoopContext) DOT() antlr.TerminalNode {
	return s.GetToken(ontologyParserDOT, 0)
}

func (s *LoopContext) Local() ILocalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Loop() (localctx ILoopContext) {
	this := p
	_ = this

	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ontologyParserRULE_loop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(393)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(394)
		p.Match(ontologyParserDOT)
	}
	{
		p.SetState(395)
		p.Local()
	}

	return localctx
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) ID() antlr.TerminalNode {
	return s.GetToken(ontologyParserID, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Label() (localctx ILabelContext) {
	this := p
	_ = this

	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ontologyParserRULE_label)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		p.Match(ontologyParserID)
	}

	return localctx
}

// ISelectContext is an interface to support dynamic dispatch.
type ISelectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectContext differentiates from other interfaces.
	IsSelectContext()
}

type SelectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectContext() *SelectContext {
	var p = new(SelectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_select
	return p
}

func (*SelectContext) IsSelectContext() {}

func NewSelectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectContext {
	var p = new(SelectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_select

	return p
}

func (s *SelectContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *SelectContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *SelectContext) BANG() antlr.TerminalNode {
	return s.GetToken(ontologyParserBANG, 0)
}

func (s *SelectContext) AllLBRA() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserLBRA)
}

func (s *SelectContext) LBRA(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserLBRA, i)
}

func (s *SelectContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserCOLON)
}

func (s *SelectContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserCOLON, i)
}

func (s *SelectContext) AllLocal() []ILocalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILocalContext); ok {
			len++
		}
	}

	tst := make([]ILocalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILocalContext); ok {
			tst[i] = t.(ILocalContext)
			i++
		}
	}

	return tst
}

func (s *SelectContext) Local(i int) ILocalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalContext)
}

func (s *SelectContext) AllRBRA() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserRBRA)
}

func (s *SelectContext) RBRA(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserRBRA, i)
}

func (s *SelectContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserOR)
}

func (s *SelectContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserOR, i)
}

func (s *SelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitSelect(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Select() (localctx ISelectContext) {
	this := p
	_ = this

	localctx = NewSelectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ontologyParserRULE_select)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(400)
		p.Match(ontologyParserBANG)
	}
	{
		p.SetState(401)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(402)
		p.Match(ontologyParserLBRA)
	}
	{
		p.SetState(403)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(404)
		p.Match(ontologyParserCOLON)
	}
	{
		p.SetState(405)
		p.Local()
	}
	{
		p.SetState(406)
		p.Match(ontologyParserRBRA)
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ontologyParserOR {
		{
			p.SetState(407)
			p.Match(ontologyParserOR)
		}
		{
			p.SetState(408)
			p.Match(ontologyParserLBRA)
		}
		{
			p.SetState(409)
			p.Match(ontologyParserID)
		}
		{
			p.SetState(410)
			p.Match(ontologyParserCOLON)
		}
		{
			p.SetState(411)
			p.Local()
		}
		{
			p.SetState(412)
			p.Match(ontologyParserRBRA)
		}

		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBranchContext is an interface to support dynamic dispatch.
type IBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBranchContext differentiates from other interfaces.
	IsBranchContext()
}

type BranchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBranchContext() *BranchContext {
	var p = new(BranchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_branch
	return p
}

func (*BranchContext) IsBranchContext() {}

func NewBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BranchContext {
	var p = new(BranchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_branch

	return p
}

func (s *BranchContext) GetParser() antlr.Parser { return s.parser }

func (s *BranchContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *BranchContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *BranchContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(ontologyParserQUESTION, 0)
}

func (s *BranchContext) AllLBRA() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserLBRA)
}

func (s *BranchContext) LBRA(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserLBRA, i)
}

func (s *BranchContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserCOLON)
}

func (s *BranchContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserCOLON, i)
}

func (s *BranchContext) AllLocal() []ILocalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILocalContext); ok {
			len++
		}
	}

	tst := make([]ILocalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILocalContext); ok {
			tst[i] = t.(ILocalContext)
			i++
		}
	}

	return tst
}

func (s *BranchContext) Local(i int) ILocalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalContext)
}

func (s *BranchContext) AllRBRA() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserRBRA)
}

func (s *BranchContext) RBRA(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserRBRA, i)
}

func (s *BranchContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserOR)
}

func (s *BranchContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserOR, i)
}

func (s *BranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BranchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitBranch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Branch() (localctx IBranchContext) {
	this := p
	_ = this

	localctx = NewBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ontologyParserRULE_branch)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(419)
		p.Match(ontologyParserQUESTION)
	}
	{
		p.SetState(420)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(421)
		p.Match(ontologyParserLBRA)
	}
	{
		p.SetState(422)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(423)
		p.Match(ontologyParserCOLON)
	}
	{
		p.SetState(424)
		p.Local()
	}
	{
		p.SetState(425)
		p.Match(ontologyParserRBRA)
	}
	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ontologyParserOR {
		{
			p.SetState(426)
			p.Match(ontologyParserOR)
		}
		{
			p.SetState(427)
			p.Match(ontologyParserLBRA)
		}
		{
			p.SetState(428)
			p.Match(ontologyParserID)
		}
		{
			p.SetState(429)
			p.Match(ontologyParserCOLON)
		}
		{
			p.SetState(430)
			p.Local()
		}
		{
			p.SetState(431)
			p.Match(ontologyParserRBRA)
		}

		p.SetState(435)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IGlobalContext is an interface to support dynamic dispatch.
type IGlobalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGlobalContext differentiates from other interfaces.
	IsGlobalContext()
}

type GlobalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobalContext() *GlobalContext {
	var p = new(GlobalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_global
	return p
}

func (*GlobalContext) IsGlobalContext() {}

func NewGlobalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalContext {
	var p = new(GlobalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_global

	return p
}

func (s *GlobalContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalContext) Pass() IPassContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPassContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPassContext)
}

func (s *GlobalContext) Global_loop() IGlobal_loopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobal_loopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobal_loopContext)
}

func (s *GlobalContext) Label() ILabelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *GlobalContext) Choice() IChoiceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChoiceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChoiceContext)
}

func (s *GlobalContext) END() antlr.TerminalNode {
	return s.GetToken(ontologyParserEND, 0)
}

func (s *GlobalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GlobalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitGlobal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Global() (localctx IGlobalContext) {
	this := p
	_ = this

	localctx = NewGlobalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ontologyParserRULE_global)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(437)
			p.Pass()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(438)
			p.Global_loop()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(439)
			p.Label()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(440)
			p.Choice()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(441)
			p.Match(ontologyParserEND)
		}

	}

	return localctx
}

// IPassContext is an interface to support dynamic dispatch.
type IPassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPassContext differentiates from other interfaces.
	IsPassContext()
}

type PassContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPassContext() *PassContext {
	var p = new(PassContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_pass
	return p
}

func (*PassContext) IsPassContext() {}

func NewPassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PassContext {
	var p = new(PassContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_pass

	return p
}

func (s *PassContext) GetParser() antlr.Parser { return s.parser }

func (s *PassContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *PassContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *PassContext) ARROW() antlr.TerminalNode {
	return s.GetToken(ontologyParserARROW, 0)
}

func (s *PassContext) COLON() antlr.TerminalNode {
	return s.GetToken(ontologyParserCOLON, 0)
}

func (s *PassContext) Type() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *PassContext) DOT() antlr.TerminalNode {
	return s.GetToken(ontologyParserDOT, 0)
}

func (s *PassContext) Global() IGlobalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobalContext)
}

func (s *PassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PassContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitPass(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Pass() (localctx IPassContext) {
	this := p
	_ = this

	localctx = NewPassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ontologyParserRULE_pass)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(445)
		p.Match(ontologyParserARROW)
	}
	{
		p.SetState(446)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(447)
		p.Match(ontologyParserCOLON)
	}
	{
		p.SetState(448)
		p.Type()
	}
	{
		p.SetState(449)
		p.Match(ontologyParserDOT)
	}
	{
		p.SetState(450)
		p.Global()
	}

	return localctx
}

// IGlobal_loopContext is an interface to support dynamic dispatch.
type IGlobal_loopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGlobal_loopContext differentiates from other interfaces.
	IsGlobal_loopContext()
}

type Global_loopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobal_loopContext() *Global_loopContext {
	var p = new(Global_loopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_global_loop
	return p
}

func (*Global_loopContext) IsGlobal_loopContext() {}

func NewGlobal_loopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Global_loopContext {
	var p = new(Global_loopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_global_loop

	return p
}

func (s *Global_loopContext) GetParser() antlr.Parser { return s.parser }

func (s *Global_loopContext) ID() antlr.TerminalNode {
	return s.GetToken(ontologyParserID, 0)
}

func (s *Global_loopContext) DOT() antlr.TerminalNode {
	return s.GetToken(ontologyParserDOT, 0)
}

func (s *Global_loopContext) Global() IGlobalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobalContext)
}

func (s *Global_loopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Global_loopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Global_loopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitGlobal_loop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Global_loop() (localctx IGlobal_loopContext) {
	this := p
	_ = this

	localctx = NewGlobal_loopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, ontologyParserRULE_global_loop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(452)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(453)
		p.Match(ontologyParserDOT)
	}
	{
		p.SetState(454)
		p.Global()
	}

	return localctx
}

// IChoiceContext is an interface to support dynamic dispatch.
type IChoiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChoiceContext differentiates from other interfaces.
	IsChoiceContext()
}

type ChoiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChoiceContext() *ChoiceContext {
	var p = new(ChoiceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_choice
	return p
}

func (*ChoiceContext) IsChoiceContext() {}

func NewChoiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChoiceContext {
	var p = new(ChoiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_choice

	return p
}

func (s *ChoiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ChoiceContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *ChoiceContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *ChoiceContext) ARROW() antlr.TerminalNode {
	return s.GetToken(ontologyParserARROW, 0)
}

func (s *ChoiceContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserCOLON)
}

func (s *ChoiceContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserCOLON, i)
}

func (s *ChoiceContext) AllLBRA() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserLBRA)
}

func (s *ChoiceContext) LBRA(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserLBRA, i)
}

func (s *ChoiceContext) AllGlobal() []IGlobalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGlobalContext); ok {
			len++
		}
	}

	tst := make([]IGlobalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGlobalContext); ok {
			tst[i] = t.(IGlobalContext)
			i++
		}
	}

	return tst
}

func (s *ChoiceContext) Global(i int) IGlobalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobalContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobalContext)
}

func (s *ChoiceContext) AllRBRA() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserRBRA)
}

func (s *ChoiceContext) RBRA(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserRBRA, i)
}

func (s *ChoiceContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserOR)
}

func (s *ChoiceContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserOR, i)
}

func (s *ChoiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChoiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChoiceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitChoice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Choice() (localctx IChoiceContext) {
	this := p
	_ = this

	localctx = NewChoiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, ontologyParserRULE_choice)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(456)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(457)
		p.Match(ontologyParserARROW)
	}
	{
		p.SetState(458)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(459)
		p.Match(ontologyParserCOLON)
	}
	{
		p.SetState(460)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(461)
		p.Match(ontologyParserLBRA)
	}
	{
		p.SetState(462)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(463)
		p.Match(ontologyParserCOLON)
	}
	{
		p.SetState(464)
		p.Global()
	}
	{
		p.SetState(465)
		p.Match(ontologyParserRBRA)
	}
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(466)
				p.Match(ontologyParserOR)
			}
			{
				p.SetState(467)
				p.Match(ontologyParserLBRA)
			}
			{
				p.SetState(468)
				p.Match(ontologyParserID)
			}
			{
				p.SetState(469)
				p.Match(ontologyParserCOLON)
			}
			{
				p.SetState(470)
				p.Global()
			}
			{
				p.SetState(471)
				p.Match(ontologyParserRBRA)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
	}

	return localctx
}

// IKnowledge_baseContext is an interface to support dynamic dispatch.
type IKnowledge_baseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKnowledge_baseContext differentiates from other interfaces.
	IsKnowledge_baseContext()
}

type Knowledge_baseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKnowledge_baseContext() *Knowledge_baseContext {
	var p = new(Knowledge_baseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_knowledge_base
	return p
}

func (*Knowledge_baseContext) IsKnowledge_baseContext() {}

func NewKnowledge_baseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Knowledge_baseContext {
	var p = new(Knowledge_baseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_knowledge_base

	return p
}

func (s *Knowledge_baseContext) GetParser() antlr.Parser { return s.parser }

func (s *Knowledge_baseContext) KNOWLEDGE() antlr.TerminalNode {
	return s.GetToken(ontologyParserKNOWLEDGE, 0)
}

func (s *Knowledge_baseContext) BASE() antlr.TerminalNode {
	return s.GetToken(ontologyParserBASE, 0)
}

func (s *Knowledge_baseContext) ID() antlr.TerminalNode {
	return s.GetToken(ontologyParserID, 0)
}

func (s *Knowledge_baseContext) LBRA() antlr.TerminalNode {
	return s.GetToken(ontologyParserLBRA, 0)
}

func (s *Knowledge_baseContext) RBRA() antlr.TerminalNode {
	return s.GetToken(ontologyParserRBRA, 0)
}

func (s *Knowledge_baseContext) AllKnowledge_base_decl() []IKnowledge_base_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKnowledge_base_declContext); ok {
			len++
		}
	}

	tst := make([]IKnowledge_base_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKnowledge_base_declContext); ok {
			tst[i] = t.(IKnowledge_base_declContext)
			i++
		}
	}

	return tst
}

func (s *Knowledge_baseContext) Knowledge_base_decl(i int) IKnowledge_base_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKnowledge_base_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKnowledge_base_declContext)
}

func (s *Knowledge_baseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Knowledge_baseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Knowledge_baseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitKnowledge_base(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Knowledge_base() (localctx IKnowledge_baseContext) {
	this := p
	_ = this

	localctx = NewKnowledge_baseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, ontologyParserRULE_knowledge_base)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(477)
		p.Match(ontologyParserKNOWLEDGE)
	}
	{
		p.SetState(478)
		p.Match(ontologyParserBASE)
	}
	{
		p.SetState(479)
		p.Match(ontologyParserID)
	}
	{
		p.SetState(480)
		p.Match(ontologyParserLBRA)
	}
	p.SetState(482)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ontologyParserESTIMATE)|(1<<ontologyParserSENSE)|(1<<ontologyParserCONTROL)|(1<<ontologyParserACTUATE))) != 0) {
		{
			p.SetState(481)
			p.Knowledge_base_decl()
		}

		p.SetState(484)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(486)
		p.Match(ontologyParserRBRA)
	}

	return localctx
}

// IKnowledge_base_declContext is an interface to support dynamic dispatch.
type IKnowledge_base_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKnowledge_base_declContext differentiates from other interfaces.
	IsKnowledge_base_declContext()
}

type Knowledge_base_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKnowledge_base_declContext() *Knowledge_base_declContext {
	var p = new(Knowledge_base_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ontologyParserRULE_knowledge_base_decl
	return p
}

func (*Knowledge_base_declContext) IsKnowledge_base_declContext() {}

func NewKnowledge_base_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Knowledge_base_declContext {
	var p = new(Knowledge_base_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ontologyParserRULE_knowledge_base_decl

	return p
}

func (s *Knowledge_base_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Knowledge_base_declContext) ESTIMATE() antlr.TerminalNode {
	return s.GetToken(ontologyParserESTIMATE, 0)
}

func (s *Knowledge_base_declContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ontologyParserID)
}

func (s *Knowledge_base_declContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ontologyParserID, i)
}

func (s *Knowledge_base_declContext) USING() antlr.TerminalNode {
	return s.GetToken(ontologyParserUSING, 0)
}

func (s *Knowledge_base_declContext) EQ() antlr.TerminalNode {
	return s.GetToken(ontologyParserEQ, 0)
}

func (s *Knowledge_base_declContext) Local() ILocalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalContext)
}

func (s *Knowledge_base_declContext) SENSE() antlr.TerminalNode {
	return s.GetToken(ontologyParserSENSE, 0)
}

func (s *Knowledge_base_declContext) CONTROL() antlr.TerminalNode {
	return s.GetToken(ontologyParserCONTROL, 0)
}

func (s *Knowledge_base_declContext) ACTUATE() antlr.TerminalNode {
	return s.GetToken(ontologyParserACTUATE, 0)
}

func (s *Knowledge_base_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Knowledge_base_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Knowledge_base_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ontologyVisitor:
		return t.VisitKnowledge_base_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ontologyParser) Knowledge_base_decl() (localctx IKnowledge_base_declContext) {
	this := p
	_ = this

	localctx = NewKnowledge_base_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, ontologyParserRULE_knowledge_base_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(512)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ontologyParserESTIMATE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(488)
			p.Match(ontologyParserESTIMATE)
		}
		{
			p.SetState(489)
			p.Match(ontologyParserID)
		}
		{
			p.SetState(490)
			p.Match(ontologyParserUSING)
		}
		{
			p.SetState(491)
			p.Match(ontologyParserID)
		}
		{
			p.SetState(492)
			p.Match(ontologyParserEQ)
		}
		{
			p.SetState(493)
			p.Local()
		}

	case ontologyParserSENSE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(494)
			p.Match(ontologyParserSENSE)
		}
		{
			p.SetState(495)
			p.Match(ontologyParserID)
		}
		{
			p.SetState(496)
			p.Match(ontologyParserUSING)
		}
		{
			p.SetState(497)
			p.Match(ontologyParserID)
		}
		{
			p.SetState(498)
			p.Match(ontologyParserEQ)
		}
		{
			p.SetState(499)
			p.Local()
		}

	case ontologyParserCONTROL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(500)
			p.Match(ontologyParserCONTROL)
		}
		{
			p.SetState(501)
			p.Match(ontologyParserID)
		}
		{
			p.SetState(502)
			p.Match(ontologyParserUSING)
		}
		{
			p.SetState(503)
			p.Match(ontologyParserID)
		}
		{
			p.SetState(504)
			p.Match(ontologyParserEQ)
		}
		{
			p.SetState(505)
			p.Local()
		}

	case ontologyParserACTUATE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(506)
			p.Match(ontologyParserACTUATE)
		}
		{
			p.SetState(507)
			p.Match(ontologyParserID)
		}
		{
			p.SetState(508)
			p.Match(ontologyParserUSING)
		}
		{
			p.SetState(509)
			p.Match(ontologyParserID)
		}
		{
			p.SetState(510)
			p.Match(ontologyParserEQ)
		}
		{
			p.SetState(511)
			p.Local()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *ontologyParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ontologyParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
