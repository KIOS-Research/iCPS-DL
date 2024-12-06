// Code generated from .\ast\parser\ontology.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ontologyLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ontologylexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ontologylexerLexerInit() {
	staticData := &ontologylexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'load'", "'show'", "'ls'", "'translate'", "'mermaid'", "'rm'",
		"'traverse'", "'configure'", "'compose'", "'project'", "'info'", "'exit'",
		"'quit'", "'domain'", "'process'", "'knowledge'", "'base'", "'property'",
		"'model'", "'class'", "'physical'", "'actuator'", "'sensor'", "'translation'",
		"'device'", "'conn'", "'estimate'", "'sense'", "'control'", "'actuate'",
		"'using'", "'local'", "'global'", "'or'", "'end'", "'('", "')'", "'{'",
		"'}'", "'['", "']'", "'@'", "'='", "':='", "','", "':'", "'.'", "'->'",
		"'!'", "'?'",
	}
	staticData.symbolicNames = []string{
		"", "LOAD", "SHOW", "LS", "TRANSLATE", "MERMAID", "REMOVE", "TRAVERSE",
		"CONFIGURE", "COMPOSE", "PROJECT", "INFO", "EXIT", "QUIT", "DOMAIN",
		"PROCESS", "KNOWLEDGE", "BASE", "PROPERTY", "MODEL", "CLASS", "PHYSICAL",
		"ACTUATOR", "SENSOR", "TRANSLATION", "DEVICE", "CONNECTION", "ESTIMATE",
		"SENSE", "CONTROL", "ACTUATE", "USING", "LOCAL", "GLOBAL", "OR", "END",
		"LPAR", "RPAR", "LBRA", "RBRA", "LSQ", "RSQ", "AT", "EQ", "ASGN", "COMMA",
		"COLON", "DOT", "ARROW", "BANG", "QUESTION", "PATH", "ID", "INT", "COMMENTS",
		"WS",
	}
	staticData.ruleNames = []string{
		"LOAD", "SHOW", "LS", "TRANSLATE", "MERMAID", "REMOVE", "TRAVERSE",
		"CONFIGURE", "COMPOSE", "PROJECT", "INFO", "EXIT", "QUIT", "DOMAIN",
		"PROCESS", "KNOWLEDGE", "BASE", "PROPERTY", "MODEL", "CLASS", "PHYSICAL",
		"ACTUATOR", "SENSOR", "TRANSLATION", "DEVICE", "CONNECTION", "ESTIMATE",
		"SENSE", "CONTROL", "ACTUATE", "USING", "LOCAL", "GLOBAL", "OR", "END",
		"LPAR", "RPAR", "LBRA", "RBRA", "LSQ", "RSQ", "AT", "EQ", "ASGN", "COMMA",
		"COLON", "DOT", "ARROW", "BANG", "QUESTION", "PATH", "ID", "INT", "COMMENTS",
		"WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 55, 421, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1,
		36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41,
		1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1,
		46, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50,
		4, 50, 388, 8, 50, 11, 50, 12, 50, 389, 1, 51, 1, 51, 5, 51, 394, 8, 51,
		10, 51, 12, 51, 397, 9, 51, 1, 52, 4, 52, 400, 8, 52, 11, 52, 12, 52, 401,
		1, 53, 1, 53, 5, 53, 406, 8, 53, 10, 53, 12, 53, 409, 9, 53, 1, 53, 3,
		53, 412, 8, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54,
		0, 0, 55, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55,
		28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73,
		37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91,
		46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54,
		109, 55, 1, 0, 5, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 1, 0, 48, 57, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32,
		425, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0,
		0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0,
		0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1,
		0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69,
		1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0,
		77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0,
		0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0,
		0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0,
		0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107,
		1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 1, 111, 1, 0, 0, 0, 3, 116, 1, 0, 0, 0,
		5, 121, 1, 0, 0, 0, 7, 124, 1, 0, 0, 0, 9, 134, 1, 0, 0, 0, 11, 142, 1,
		0, 0, 0, 13, 145, 1, 0, 0, 0, 15, 154, 1, 0, 0, 0, 17, 164, 1, 0, 0, 0,
		19, 172, 1, 0, 0, 0, 21, 180, 1, 0, 0, 0, 23, 185, 1, 0, 0, 0, 25, 190,
		1, 0, 0, 0, 27, 195, 1, 0, 0, 0, 29, 202, 1, 0, 0, 0, 31, 210, 1, 0, 0,
		0, 33, 220, 1, 0, 0, 0, 35, 225, 1, 0, 0, 0, 37, 234, 1, 0, 0, 0, 39, 240,
		1, 0, 0, 0, 41, 246, 1, 0, 0, 0, 43, 255, 1, 0, 0, 0, 45, 264, 1, 0, 0,
		0, 47, 271, 1, 0, 0, 0, 49, 283, 1, 0, 0, 0, 51, 290, 1, 0, 0, 0, 53, 295,
		1, 0, 0, 0, 55, 304, 1, 0, 0, 0, 57, 310, 1, 0, 0, 0, 59, 318, 1, 0, 0,
		0, 61, 326, 1, 0, 0, 0, 63, 332, 1, 0, 0, 0, 65, 338, 1, 0, 0, 0, 67, 345,
		1, 0, 0, 0, 69, 348, 1, 0, 0, 0, 71, 352, 1, 0, 0, 0, 73, 354, 1, 0, 0,
		0, 75, 356, 1, 0, 0, 0, 77, 358, 1, 0, 0, 0, 79, 360, 1, 0, 0, 0, 81, 362,
		1, 0, 0, 0, 83, 364, 1, 0, 0, 0, 85, 366, 1, 0, 0, 0, 87, 368, 1, 0, 0,
		0, 89, 371, 1, 0, 0, 0, 91, 373, 1, 0, 0, 0, 93, 375, 1, 0, 0, 0, 95, 377,
		1, 0, 0, 0, 97, 380, 1, 0, 0, 0, 99, 382, 1, 0, 0, 0, 101, 384, 1, 0, 0,
		0, 103, 391, 1, 0, 0, 0, 105, 399, 1, 0, 0, 0, 107, 403, 1, 0, 0, 0, 109,
		417, 1, 0, 0, 0, 111, 112, 5, 108, 0, 0, 112, 113, 5, 111, 0, 0, 113, 114,
		5, 97, 0, 0, 114, 115, 5, 100, 0, 0, 115, 2, 1, 0, 0, 0, 116, 117, 5, 115,
		0, 0, 117, 118, 5, 104, 0, 0, 118, 119, 5, 111, 0, 0, 119, 120, 5, 119,
		0, 0, 120, 4, 1, 0, 0, 0, 121, 122, 5, 108, 0, 0, 122, 123, 5, 115, 0,
		0, 123, 6, 1, 0, 0, 0, 124, 125, 5, 116, 0, 0, 125, 126, 5, 114, 0, 0,
		126, 127, 5, 97, 0, 0, 127, 128, 5, 110, 0, 0, 128, 129, 5, 115, 0, 0,
		129, 130, 5, 108, 0, 0, 130, 131, 5, 97, 0, 0, 131, 132, 5, 116, 0, 0,
		132, 133, 5, 101, 0, 0, 133, 8, 1, 0, 0, 0, 134, 135, 5, 109, 0, 0, 135,
		136, 5, 101, 0, 0, 136, 137, 5, 114, 0, 0, 137, 138, 5, 109, 0, 0, 138,
		139, 5, 97, 0, 0, 139, 140, 5, 105, 0, 0, 140, 141, 5, 100, 0, 0, 141,
		10, 1, 0, 0, 0, 142, 143, 5, 114, 0, 0, 143, 144, 5, 109, 0, 0, 144, 12,
		1, 0, 0, 0, 145, 146, 5, 116, 0, 0, 146, 147, 5, 114, 0, 0, 147, 148, 5,
		97, 0, 0, 148, 149, 5, 118, 0, 0, 149, 150, 5, 101, 0, 0, 150, 151, 5,
		114, 0, 0, 151, 152, 5, 115, 0, 0, 152, 153, 5, 101, 0, 0, 153, 14, 1,
		0, 0, 0, 154, 155, 5, 99, 0, 0, 155, 156, 5, 111, 0, 0, 156, 157, 5, 110,
		0, 0, 157, 158, 5, 102, 0, 0, 158, 159, 5, 105, 0, 0, 159, 160, 5, 103,
		0, 0, 160, 161, 5, 117, 0, 0, 161, 162, 5, 114, 0, 0, 162, 163, 5, 101,
		0, 0, 163, 16, 1, 0, 0, 0, 164, 165, 5, 99, 0, 0, 165, 166, 5, 111, 0,
		0, 166, 167, 5, 109, 0, 0, 167, 168, 5, 112, 0, 0, 168, 169, 5, 111, 0,
		0, 169, 170, 5, 115, 0, 0, 170, 171, 5, 101, 0, 0, 171, 18, 1, 0, 0, 0,
		172, 173, 5, 112, 0, 0, 173, 174, 5, 114, 0, 0, 174, 175, 5, 111, 0, 0,
		175, 176, 5, 106, 0, 0, 176, 177, 5, 101, 0, 0, 177, 178, 5, 99, 0, 0,
		178, 179, 5, 116, 0, 0, 179, 20, 1, 0, 0, 0, 180, 181, 5, 105, 0, 0, 181,
		182, 5, 110, 0, 0, 182, 183, 5, 102, 0, 0, 183, 184, 5, 111, 0, 0, 184,
		22, 1, 0, 0, 0, 185, 186, 5, 101, 0, 0, 186, 187, 5, 120, 0, 0, 187, 188,
		5, 105, 0, 0, 188, 189, 5, 116, 0, 0, 189, 24, 1, 0, 0, 0, 190, 191, 5,
		113, 0, 0, 191, 192, 5, 117, 0, 0, 192, 193, 5, 105, 0, 0, 193, 194, 5,
		116, 0, 0, 194, 26, 1, 0, 0, 0, 195, 196, 5, 100, 0, 0, 196, 197, 5, 111,
		0, 0, 197, 198, 5, 109, 0, 0, 198, 199, 5, 97, 0, 0, 199, 200, 5, 105,
		0, 0, 200, 201, 5, 110, 0, 0, 201, 28, 1, 0, 0, 0, 202, 203, 5, 112, 0,
		0, 203, 204, 5, 114, 0, 0, 204, 205, 5, 111, 0, 0, 205, 206, 5, 99, 0,
		0, 206, 207, 5, 101, 0, 0, 207, 208, 5, 115, 0, 0, 208, 209, 5, 115, 0,
		0, 209, 30, 1, 0, 0, 0, 210, 211, 5, 107, 0, 0, 211, 212, 5, 110, 0, 0,
		212, 213, 5, 111, 0, 0, 213, 214, 5, 119, 0, 0, 214, 215, 5, 108, 0, 0,
		215, 216, 5, 101, 0, 0, 216, 217, 5, 100, 0, 0, 217, 218, 5, 103, 0, 0,
		218, 219, 5, 101, 0, 0, 219, 32, 1, 0, 0, 0, 220, 221, 5, 98, 0, 0, 221,
		222, 5, 97, 0, 0, 222, 223, 5, 115, 0, 0, 223, 224, 5, 101, 0, 0, 224,
		34, 1, 0, 0, 0, 225, 226, 5, 112, 0, 0, 226, 227, 5, 114, 0, 0, 227, 228,
		5, 111, 0, 0, 228, 229, 5, 112, 0, 0, 229, 230, 5, 101, 0, 0, 230, 231,
		5, 114, 0, 0, 231, 232, 5, 116, 0, 0, 232, 233, 5, 121, 0, 0, 233, 36,
		1, 0, 0, 0, 234, 235, 5, 109, 0, 0, 235, 236, 5, 111, 0, 0, 236, 237, 5,
		100, 0, 0, 237, 238, 5, 101, 0, 0, 238, 239, 5, 108, 0, 0, 239, 38, 1,
		0, 0, 0, 240, 241, 5, 99, 0, 0, 241, 242, 5, 108, 0, 0, 242, 243, 5, 97,
		0, 0, 243, 244, 5, 115, 0, 0, 244, 245, 5, 115, 0, 0, 245, 40, 1, 0, 0,
		0, 246, 247, 5, 112, 0, 0, 247, 248, 5, 104, 0, 0, 248, 249, 5, 121, 0,
		0, 249, 250, 5, 115, 0, 0, 250, 251, 5, 105, 0, 0, 251, 252, 5, 99, 0,
		0, 252, 253, 5, 97, 0, 0, 253, 254, 5, 108, 0, 0, 254, 42, 1, 0, 0, 0,
		255, 256, 5, 97, 0, 0, 256, 257, 5, 99, 0, 0, 257, 258, 5, 116, 0, 0, 258,
		259, 5, 117, 0, 0, 259, 260, 5, 97, 0, 0, 260, 261, 5, 116, 0, 0, 261,
		262, 5, 111, 0, 0, 262, 263, 5, 114, 0, 0, 263, 44, 1, 0, 0, 0, 264, 265,
		5, 115, 0, 0, 265, 266, 5, 101, 0, 0, 266, 267, 5, 110, 0, 0, 267, 268,
		5, 115, 0, 0, 268, 269, 5, 111, 0, 0, 269, 270, 5, 114, 0, 0, 270, 46,
		1, 0, 0, 0, 271, 272, 5, 116, 0, 0, 272, 273, 5, 114, 0, 0, 273, 274, 5,
		97, 0, 0, 274, 275, 5, 110, 0, 0, 275, 276, 5, 115, 0, 0, 276, 277, 5,
		108, 0, 0, 277, 278, 5, 97, 0, 0, 278, 279, 5, 116, 0, 0, 279, 280, 5,
		105, 0, 0, 280, 281, 5, 111, 0, 0, 281, 282, 5, 110, 0, 0, 282, 48, 1,
		0, 0, 0, 283, 284, 5, 100, 0, 0, 284, 285, 5, 101, 0, 0, 285, 286, 5, 118,
		0, 0, 286, 287, 5, 105, 0, 0, 287, 288, 5, 99, 0, 0, 288, 289, 5, 101,
		0, 0, 289, 50, 1, 0, 0, 0, 290, 291, 5, 99, 0, 0, 291, 292, 5, 111, 0,
		0, 292, 293, 5, 110, 0, 0, 293, 294, 5, 110, 0, 0, 294, 52, 1, 0, 0, 0,
		295, 296, 5, 101, 0, 0, 296, 297, 5, 115, 0, 0, 297, 298, 5, 116, 0, 0,
		298, 299, 5, 105, 0, 0, 299, 300, 5, 109, 0, 0, 300, 301, 5, 97, 0, 0,
		301, 302, 5, 116, 0, 0, 302, 303, 5, 101, 0, 0, 303, 54, 1, 0, 0, 0, 304,
		305, 5, 115, 0, 0, 305, 306, 5, 101, 0, 0, 306, 307, 5, 110, 0, 0, 307,
		308, 5, 115, 0, 0, 308, 309, 5, 101, 0, 0, 309, 56, 1, 0, 0, 0, 310, 311,
		5, 99, 0, 0, 311, 312, 5, 111, 0, 0, 312, 313, 5, 110, 0, 0, 313, 314,
		5, 116, 0, 0, 314, 315, 5, 114, 0, 0, 315, 316, 5, 111, 0, 0, 316, 317,
		5, 108, 0, 0, 317, 58, 1, 0, 0, 0, 318, 319, 5, 97, 0, 0, 319, 320, 5,
		99, 0, 0, 320, 321, 5, 116, 0, 0, 321, 322, 5, 117, 0, 0, 322, 323, 5,
		97, 0, 0, 323, 324, 5, 116, 0, 0, 324, 325, 5, 101, 0, 0, 325, 60, 1, 0,
		0, 0, 326, 327, 5, 117, 0, 0, 327, 328, 5, 115, 0, 0, 328, 329, 5, 105,
		0, 0, 329, 330, 5, 110, 0, 0, 330, 331, 5, 103, 0, 0, 331, 62, 1, 0, 0,
		0, 332, 333, 5, 108, 0, 0, 333, 334, 5, 111, 0, 0, 334, 335, 5, 99, 0,
		0, 335, 336, 5, 97, 0, 0, 336, 337, 5, 108, 0, 0, 337, 64, 1, 0, 0, 0,
		338, 339, 5, 103, 0, 0, 339, 340, 5, 108, 0, 0, 340, 341, 5, 111, 0, 0,
		341, 342, 5, 98, 0, 0, 342, 343, 5, 97, 0, 0, 343, 344, 5, 108, 0, 0, 344,
		66, 1, 0, 0, 0, 345, 346, 5, 111, 0, 0, 346, 347, 5, 114, 0, 0, 347, 68,
		1, 0, 0, 0, 348, 349, 5, 101, 0, 0, 349, 350, 5, 110, 0, 0, 350, 351, 5,
		100, 0, 0, 351, 70, 1, 0, 0, 0, 352, 353, 5, 40, 0, 0, 353, 72, 1, 0, 0,
		0, 354, 355, 5, 41, 0, 0, 355, 74, 1, 0, 0, 0, 356, 357, 5, 123, 0, 0,
		357, 76, 1, 0, 0, 0, 358, 359, 5, 125, 0, 0, 359, 78, 1, 0, 0, 0, 360,
		361, 5, 91, 0, 0, 361, 80, 1, 0, 0, 0, 362, 363, 5, 93, 0, 0, 363, 82,
		1, 0, 0, 0, 364, 365, 5, 64, 0, 0, 365, 84, 1, 0, 0, 0, 366, 367, 5, 61,
		0, 0, 367, 86, 1, 0, 0, 0, 368, 369, 5, 58, 0, 0, 369, 370, 5, 61, 0, 0,
		370, 88, 1, 0, 0, 0, 371, 372, 5, 44, 0, 0, 372, 90, 1, 0, 0, 0, 373, 374,
		5, 58, 0, 0, 374, 92, 1, 0, 0, 0, 375, 376, 5, 46, 0, 0, 376, 94, 1, 0,
		0, 0, 377, 378, 5, 45, 0, 0, 378, 379, 5, 62, 0, 0, 379, 96, 1, 0, 0, 0,
		380, 381, 5, 33, 0, 0, 381, 98, 1, 0, 0, 0, 382, 383, 5, 63, 0, 0, 383,
		100, 1, 0, 0, 0, 384, 387, 3, 103, 51, 0, 385, 386, 5, 47, 0, 0, 386, 388,
		3, 103, 51, 0, 387, 385, 1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389, 387, 1,
		0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 102, 1, 0, 0, 0, 391, 395, 7, 0, 0,
		0, 392, 394, 7, 1, 0, 0, 393, 392, 1, 0, 0, 0, 394, 397, 1, 0, 0, 0, 395,
		393, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 104, 1, 0, 0, 0, 397, 395,
		1, 0, 0, 0, 398, 400, 7, 2, 0, 0, 399, 398, 1, 0, 0, 0, 400, 401, 1, 0,
		0, 0, 401, 399, 1, 0, 0, 0, 401, 402, 1, 0, 0, 0, 402, 106, 1, 0, 0, 0,
		403, 407, 5, 35, 0, 0, 404, 406, 8, 3, 0, 0, 405, 404, 1, 0, 0, 0, 406,
		409, 1, 0, 0, 0, 407, 405, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408, 411,
		1, 0, 0, 0, 409, 407, 1, 0, 0, 0, 410, 412, 5, 13, 0, 0, 411, 410, 1, 0,
		0, 0, 411, 412, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 414, 5, 10, 0, 0,
		414, 415, 1, 0, 0, 0, 415, 416, 6, 53, 0, 0, 416, 108, 1, 0, 0, 0, 417,
		418, 7, 4, 0, 0, 418, 419, 1, 0, 0, 0, 419, 420, 6, 54, 0, 0, 420, 110,
		1, 0, 0, 0, 6, 0, 389, 395, 401, 407, 411, 1, 6, 0, 0,
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

// ontologyLexerInit initializes any static state used to implement ontologyLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewontologyLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func OntologyLexerInit() {
	staticData := &ontologylexerLexerStaticData
	staticData.once.Do(ontologylexerLexerInit)
}

// NewontologyLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewontologyLexer(input antlr.CharStream) *ontologyLexer {
	OntologyLexerInit()
	l := new(ontologyLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ontologylexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "ontology.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ontologyLexer tokens.
const (
	ontologyLexerLOAD        = 1
	ontologyLexerSHOW        = 2
	ontologyLexerLS          = 3
	ontologyLexerTRANSLATE   = 4
	ontologyLexerMERMAID     = 5
	ontologyLexerREMOVE      = 6
	ontologyLexerTRAVERSE    = 7
	ontologyLexerCONFIGURE   = 8
	ontologyLexerCOMPOSE     = 9
	ontologyLexerPROJECT     = 10
	ontologyLexerINFO        = 11
	ontologyLexerEXIT        = 12
	ontologyLexerQUIT        = 13
	ontologyLexerDOMAIN      = 14
	ontologyLexerPROCESS     = 15
	ontologyLexerKNOWLEDGE   = 16
	ontologyLexerBASE        = 17
	ontologyLexerPROPERTY    = 18
	ontologyLexerMODEL       = 19
	ontologyLexerCLASS       = 20
	ontologyLexerPHYSICAL    = 21
	ontologyLexerACTUATOR    = 22
	ontologyLexerSENSOR      = 23
	ontologyLexerTRANSLATION = 24
	ontologyLexerDEVICE      = 25
	ontologyLexerCONNECTION  = 26
	ontologyLexerESTIMATE    = 27
	ontologyLexerSENSE       = 28
	ontologyLexerCONTROL     = 29
	ontologyLexerACTUATE     = 30
	ontologyLexerUSING       = 31
	ontologyLexerLOCAL       = 32
	ontologyLexerGLOBAL      = 33
	ontologyLexerOR          = 34
	ontologyLexerEND         = 35
	ontologyLexerLPAR        = 36
	ontologyLexerRPAR        = 37
	ontologyLexerLBRA        = 38
	ontologyLexerRBRA        = 39
	ontologyLexerLSQ         = 40
	ontologyLexerRSQ         = 41
	ontologyLexerAT          = 42
	ontologyLexerEQ          = 43
	ontologyLexerASGN        = 44
	ontologyLexerCOMMA       = 45
	ontologyLexerCOLON       = 46
	ontologyLexerDOT         = 47
	ontologyLexerARROW       = 48
	ontologyLexerBANG        = 49
	ontologyLexerQUESTION    = 50
	ontologyLexerPATH        = 51
	ontologyLexerID          = 52
	ontologyLexerINT         = 53
	ontologyLexerCOMMENTS    = 54
	ontologyLexerWS          = 55
)
