grammar R2D2;

// Lexer Rules
// Comentários
COMMENT: '//' ~[\r\n]* -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;

// Palavras-chave
INTERFACE: 'interface';
MODULE: 'module';
IMPLEMENTS: 'implements';
IMPORT: 'import';
FROM: 'from';
EXPORT: 'export';
FUNC: 'fn';
PSEUDO: 'pseudo';
IF: 'if';
ELSE: 'else';
FOR: 'for';
WHILE: 'while';
LOOP: 'loop';
BREAK: 'break';
RETURN: 'return';
SEND: 'send';
VAR: 'var';
LET: 'let';
CONST: 'const';

// Tipos Primitivos
TYPE: (
    'i8' | 'i16' | 'i32' | 'i64' |    // Inteiros com sinal
    'u8' | 'u16' | 'u32' | 'u64' |    // Inteiros sem sinal
    'f32' | 'f64' |                   // Ponto flutuante
    'bool' | 'string' | 'void'        // Outros tipos básicos
);

// Operadores básicos (C-like)
ASSIGN: '=';
PLUS: '+';
MINUS: '-';
MULT: '*';
DIV: '/';
MOD: '%';
INCREMENT: '++';
DECREMENT: '--';
EQ: '==';
NEQ: '!=';
LT: '<';
GT: '>';
LEQ: '<=';
GEQ: '>=';
AND: '&&';
OR: '||';
NOT: '!';

// Pontuação
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
LBRACK: '[';
RBRACK: ']';
COMMA: ',';
DOT: '.';
COLON: ':';
SEMI: ';';

// Identificadores e literais
IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;
INTEGER: [0-9]+;
FLOAT: [0-9]+ '.' [0-9]+;
STRING_LITERAL: '"' (~["\r\n\\] | '\\' .)* '"';
BOOL_LITERAL: 'true' | 'false';

// Ignorar whitespace
WS: [ \t\r\n]+ -> skip;

// Parser Rules
program
    : (importDeclaration | interfaceDeclaration | moduleDeclaration | functionDeclaration)*
    ;

importDeclaration
    : IMPORT STRING_LITERAL FROM STRING_LITERAL SEMI
    ;

interfaceDeclaration
    : INTERFACE IDENTIFIER LBRACE
        (exportFunctionDeclaration)*
      RBRACE
    ;

moduleDeclaration
    : MODULE IDENTIFIER (IMPLEMENTS IDENTIFIER)? LBRACE
        (variableDeclaration | functionDeclaration)*
      RBRACE
    ;

exportFunctionDeclaration
    : EXPORT FUNC IDENTIFIER LPAREN parameterList? RPAREN (COLON type)? SEMI
    ;

functionDeclaration
    : (EXPORT | PSEUDO)? FUNC IDENTIFIER LPAREN parameterList? RPAREN (COLON type)?
      (block | SEMI)
    ;

parameterList
    : parameter (COMMA parameter)*
    ;

parameter
    : IDENTIFIER COLON type
    ;

type
    : TYPE                         // Tipos primitivos
    | IDENTIFIER                   // Tipos definidos pelo usuário
    | type LBRACK INTEGER? RBRACK  // Arrays
    ;

variableDeclaration
    : (VAR | LET | CONST) IDENTIFIER (COLON type)? ASSIGN expression SEMI
    ;

statement
    : variableDeclaration
    | expressionStatement
    | ifStatement
    | forStatement
    | whileStatement
    | loopStatement
    | breakStatement
    | returnStatement
    | block
    ;

expressionStatement
    : expression SEMI
    ;

ifStatement
    : IF LPAREN expression RPAREN 
      block
      (ELSE IF LPAREN expression RPAREN block)*
      (ELSE block)?
    ;

forStatement
    : FOR (IDENTIFIER COLON)? type? IDENTIFIER? ASSIGN expression SEMI 
          expression SEMI 
          (expression | IDENTIFIER (INCREMENT | DECREMENT))
      block
    ;

whileStatement
    : WHILE LPAREN expression RPAREN block
    ;

loopStatement
    : LOOP block
    ;

breakStatement
    : BREAK SEMI
    ;

returnStatement
    : RETURN expression? SEMI
    ;

expression
    : primary
    | IDENTIFIER
    | functionCall
    | IDENTIFIER DOT IDENTIFIER            // Acesso a membros
    | IDENTIFIER (INCREMENT | DECREMENT)   // Pós incremento/decremento
    | (INCREMENT | DECREMENT) IDENTIFIER   // Pré incremento/decremento
    | LPAREN expression RPAREN
    | expression (MULT | DIV | MOD) expression
    | expression (PLUS | MINUS) expression
    | expression (LT | GT | LEQ | GEQ) expression
    | expression (EQ | NEQ) expression
    | expression AND expression
    | expression OR expression
    | NOT expression
    | SEND expression
    ;

primary
    : INTEGER
    | FLOAT
    | STRING_LITERAL
    | BOOL_LITERAL
    ;

functionCall
    : IDENTIFIER LPAREN (expression (COMMA expression)*)? RPAREN
    ;

block
    : LBRACE statement* RBRACE
    ;
