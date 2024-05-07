grammar ontology;


program         :   expression* exit? EOF
                ;

exit            :   EXIT
                |   QUIT
                ;

expression      :   paradigm
                |   process
                |   local_configuration
                |   global_configuration
                |   knowledge_base
                |   load
                |   assignment
                |   remove
                |   show
                |   mermaid
                |   translate
                |   traverse
                |   configure
                |   compose
                |   project
                |   info
                |   render
                |   expression LSQ INT RSQ
                |   ID
                |   LPAR expression RPAR
                ;

load            :   LOAD (ID | PATH)
                ;

assignment      :   ID ASGN expression
                ;

remove          :   REMOVE ID
                ;

show            :   SHOW expression?
                |   LS expression?
                ;

mermaid         :   MERMAID expression
                ;

translate       :   TRANSLATE expression
                ;

traverse        :   TRAVERSE ID DOT ID expression
                ;

configure       :   CONFIGURE expression expression ID ID
                ;

compose         :   COMPOSE expression
                ;

project         :   PROJECT expression
                ;

info            :   INFO expression?
                ;

render          :   RENDER expression
                ;

paradigm        :    PARADIGM LBRA paradigm_declaration* RBRA
                ;

paradigm_declaration
                :   property
                |   class
                |   model
                |   translation
                ;

property        :   PROPERTY type (COMMA type)*
                ;

type            :   ID LBRA ID (COMMA ID)* RBRA
                |   ID
                ;

model           :   MODEL ID (COMMA ID)*
                ;

class           :   (PHYSICAL | ACTUATOR) ID LPAR (ID (COMMA ID)*)? RPAR COLON internal_edge (COMMA internal_edge)
                ;

internal_edge   :   ID ARROW ID
                ;

translation     :   TRANSLATION ID ARROW ID COLON arg_connection (COMMA arg_connection)*
                ;

arg_connection  :   ID DOT ID ARROW ID DOT ID //qualified
                ;

connection      :   ID ARROW ID
                ;

process         :   PROCESS ID LBRA process_declaration* RBRA
                ;

process_declaration
                :   device
                |   component
                |   connection_decl
//                |   agent
                ;

device          :   DEVICE ID (COMMA ID)*
                ;

component       :   PHYSICAL ID (COMMA ID)* ID
                |   ACTUATOR ID AT ID (COMMA ID AT ID)* ID
                |   SENSOR ID AT ID (COMMA ID AT ID)* ID
                ;

connection_decl :   CONNECTION connection (COMMA connection)*
                ;

local_configuration
                :   LOCAL LBRA (ID EQ local)+ RBRA
                ;

global_configuration
                :   GLOBAL global
                ;

local           :   END
                |   send
                |   receive
                |   loop
                |   label
                |   select
                |   branch
                ;

send            :   ID BANG type DOT local
                ;

receive         :   ID QUESTION type DOT local
                ;

loop            :   ID DOT local
                ;

label           :   ID
                ;

select          :   ID BANG ID LBRA ID COLON local RBRA (OR LBRA ID COLON local RBRA)+
                ;

branch          :   ID QUESTION ID LBRA ID COLON local RBRA (OR LBRA ID COLON local RBRA)+
                ;

global          :   pass
                |   global_loop
                |   label
                |   choice
                |   END
                ;

pass            :   ID ARROW ID COLON type DOT global
                ;

global_loop     :   ID DOT global
                ;

choice          :   ID ARROW ID COLON ID LBRA ID COLON global RBRA (OR LBRA ID COLON global RBRA)+
                ;


knowledge_base  :   KNOWLEDGE BASE ID LBRA knowledge_base_decl+ RBRA
                ;

knowledge_base_decl
                :   ESTIMATE ID USING ID EQ local
                |   SENSE ID USING ID EQ local
                |   CONTROL ID USING ID EQ local
                |   ACTUATE ID USING ID EQ local
                ;


LOAD:               'load';
SHOW:               'show';
LS:                 'ls';
TRANSLATE:          'translate';
MERMAID:            'mermaid';
REMOVE:             'rm';
TRAVERSE:           'traverse';
CONFIGURE:          'configure';
COMPOSE:            'compose';
PROJECT:            'project';
INFO:               'info';
RENDER:             'render';
EXIT:               'exit';
QUIT:               'quit';

PARADIGM:           'paradigm';
PROCESS:            'process';
KNOWLEDGE:          'knowledge';
BASE:               'base';

PROPERTY:           'property';
MODEL:              'model';
CLASS:              'class';
PHYSICAL:           'physical';
ACTUATOR:           'actuator';
SENSOR:             'sensor';
TRANSLATION:        'translation';

DEVICE:             'device';
CONNECTION:         'conn';

ESTIMATE:           'estimate';
SENSE:              'sense';
CONTROL:            'control';
ACTUATE:            'actuate';
USING:              'using';

LOCAL:              'local';
GLOBAL:             'global';

//CHOICE:             'choice';
OR:                 'or';
END:                'end';

LPAR:               '(';
RPAR:               ')';
LBRA:               '{';
RBRA:               '}';
LSQ:                '[';
RSQ:                ']';
AT:                 '@';
EQ:                 '=';
ASGN:               ':=';
COMMA:              ',';
COLON:              ':';
DOT:                '.';
ARROW:              '->';
BANG:               '!';
QUESTION:           '?';

//NODE:               ID DOT ID;
PATH:               ID ('/' ID)+;
ID:                 [a-zA-Z][a-zA-Z0-9_]*;
INT:                [0-9]+;

COMMENTS:   '#' ~( '\n'|'\r' )* '\r'? '\n'  -> skip;
WS:         [ \t\n\r] -> skip;
