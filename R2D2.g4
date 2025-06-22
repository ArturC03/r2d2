grammar R2D2;

/*
 * Parser Rules
 */

program
  : importDeclaration* interfaceDeclaration* declaration* EOF
  ;

declaration
  : moduleDeclaration
  | globalDeclaration
  | typeDeclaration
  ;

globalDeclaration
  : CONST IDENTIFIER typeExpression ASSIGN expression SEMI
  ;

importDeclaration
  : USE STRING_LITERAL SEMI
  ;

interfaceDeclaration
  : INTERFACE IDENTIFIER (IMPLEMENTS IDENTIFIER)? LBRACE (functionDeclaration | variableDeclaration)* RBRACE
  ;

moduleDeclaration
  : MODULE IDENTIFIER (IMPLEMENTS IDENTIFIER)? LBRACE (functionDeclaration | typeDeclaration | variableDeclaration)* RBRACE
  ;

functionDeclaration
  : (EXPORT)? (PSEUDO)? FN IDENTIFIER LPAREN parameterList? RPAREN (typeExpression)? (block | SEMI)
  ;

functionCallStatement
  : functionCall SEMI
  ;

functionCall
  : (IDENTIFIER DOT)* IDENTIFIER LPAREN argumentList? RPAREN
  ;

parameterList
  : parameter (COMMA parameter)*
  ;

parameter
  : IDENTIFIER ( typeExpression )?
  ;

typeExpression
  : baseType arrayDimensions?
  | arrayDimensions? baseType
  ;

arrayDimensions
  : (LBRACK (INT_LITERAL)? RBRACK)+
  ;

baseType
  : IDENTIFIER
  | TYPE
  | genericType
  ;

genericType
  : IDENTIFIER LT typeExpression (COMMA typeExpression)* GT
  ;

typeDeclaration
  : (EXPORT)? 'type' IDENTIFIER LBRACE (variableDeclaration)* RBRACE
  ;

variableDeclaration
  : (EXPORT)? (VAR | LET | CONST) IDENTIFIER (typeExpression)? (ASSIGN expression)? SEMI
  ;

statement
  : variableDeclaration
  | functionCallStatement
  | expressionStatement
  | ifStatement
  | forStatement
  | whileStatement
  | loopStatement
  | cicleControl
  | returnStatement
  | switchStatement
  | assignmentDeclaration
  | jsStatement
  ;

expressionStatement
  : expression SEMI
  ;

ifStatement
  : IF (LPAREN)? expression (RPAREN)? (block | ARROW statement)
    (ELSE IF (LPAREN)? expression (RPAREN)? (block | ARROW statement))*
    (ELSE (block | ARROW statement))?
  ;

forStatement
  : FOR (LPAREN)? simpleFor (RPAREN)? block
  ;

assignmentDeclaration
  : assignment SEMI
  ;

assignment
  : IDENTIFIER assignmentOperator expression
  | IDENTIFIER (INCREMENT | DECREMENT)
  | IDENTIFIER LBRACK expression RBRACK assignmentOperator expression
  | IDENTIFIER LBRACK expression RBRACK (INCREMENT | DECREMENT)
  ;

assignmentOperator
  : ASSIGN
  | PLUS_ASSIGN
  | MINUS_ASSIGN
  | MULT_ASSIGN
  | DIV_ASSIGN
  | MOD_ASSIGN
  ;

simpleFor
  : (variableDeclaration | assignment SEMI)? (expression SEMI)? (assignment)?
  ;

whileStatement
  : WHILE (LPAREN)? expression (RPAREN)? block
  ;

loopStatement
  : LOOP block
  ;

cicleControl
  : (breakStatement | continueStatement)
  ;

breakStatement
  : BREAK SEMI
  ;

continueStatement
  : CONTINUE SEMI
  ;

returnStatement
  : RETURN expression? SEMI
  ;

expression 
  : literal                                                #literalExpression
  | IDENTIFIER                                             #identifierExpression
  | functionCall                                           #functionCallExpression
  | expression LBRACK expression RBRACK                    #arrayAccessExpression
  | '(' expression ')'                                     #parenthesisExpression
  | (NOT | MINUS | INCREMENT | DECREMENT) expression       #unaryExpression
  | expression (MULT | DIV | MOD) expression               #multiplicativeExpression
  | expression (PLUS | MINUS) expression                   #additiveExpression
  | expression (EQ | NEQ | LT | GT | LEQ | GEQ) expression #comparisonExpression
  | expression (AND | OR) expression                       #logicalExpression
  ;

argumentList
  : expression (COMMA expression)*
  ;

objectLiteral
  : LBRACK ((expression ARROW expression)(COMMA (expression ARROW expression))*)? RBRACK
  ;

// Array literals with comma-separated values
arrayLiteral
  : LBRACK (expression (COMMA expression)*)? RBRACK
  ;

literal
  : objectLiteral
  | arrayLiteral
  | INT_LITERAL
  | FLOAT_LITERAL
  | STRING_LITERAL
  | BOOL_LITERAL
  | NULL_LITERAL
  ;

block
  : LBRACE statement* RBRACE
  ;

switchStatement
  : SWITCH (LPAREN)? expression (RPAREN)? LBRACE switchCase* defaultCase? RBRACE
  ;

switchCase
  : CASE expression block
  | CASE expression ARROW statement
  ;

defaultCase
  : DEFAULT block
  | DEFAULT ARROW statement
  ;

jsStatement
  : AT JS STRING_LITERAL SEMI
  ;

/*
 * Lexer Rules
 */

// Keywords
USE: 'use';
IMPORT: 'import';
FROM: 'from';
INTERFACE: 'interface';
MODULE: 'module';
IMPLEMENTS: '::';
EXPORT: 'export';
FN: 'fn';
PSEUDO: 'pseudo';
VAR: 'var';
LET: 'let';
CONST: 'const';
IF: 'if';
ELSE: 'else';
LOOP: 'loop';
FOR: 'for';
WHILE: 'while';
BREAK: 'break';
SEND: 'send';
CONTINUE: 'continue';
RETURN: 'return';
SWITCH: 'switch';
CASE: 'case';
DEFAULT: 'default';

// Operators (order matters - longer operators first!)
ARROW: '=>';  // Used for associative arrays, must be before other operators that might match '=' or '>'
INCREMENT: '++';
DECREMENT: '--';
PLUS_ASSIGN: '+=';
MINUS_ASSIGN: '-=';
MULT_ASSIGN: '*=';
DIV_ASSIGN: '/=';
MOD_ASSIGN: '%=';
EQ: '==';
NEQ: '!=';
LEQ: '<=';
GEQ: '>=';
AND: '&&';
OR: '||';
PLUS: '+';
MINUS: '-';
MULT: '*';
DIV: '/';
MOD: '%';
ASSIGN: '=';
LT: '<';
GT: '>';
NOT: '!';

// Delimiters
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

// Other stuff
AT      : '@';
JS      : 'js';

TYPE
  : 'number'
  | 'boolean'
  | 'string'
  | 'array'
  | 'object'
  | 'void'
  ;

STRING_LITERAL
  : '"""' .*? '"""'              // Multilinha — permissiva e permite \n
  | '"' ( '\\' . | ~["\\\r\n] )* '"'  // Linha única com escapes
  ;

fragment ESCAPE_SEQUENCE : '\\' [btnr"\\];

BOOL_LITERAL
  : 'true'
  | 'false'
  ;

NULL_LITERAL
  : 'null'
  ;

// Fixed: Removed SignPart from INT_LITERAL
INT_LITERAL
  : DecimalIntegerLiteral
  | HexIntegerLiteral
  | OctalIntegerLiteral
  | BinaryIntegerLiteral
  ;

// Fixed: Removed SignPart from FLOAT_LITERAL
FLOAT_LITERAL
  : DecimalNumeral '.' DecimalDigits? ExponentPart?
  | '.' DecimalDigits ExponentPart?
  | DecimalNumeral ExponentPart
  ;

// Identifiers and literals
IDENTIFIER: ([a-zA-Z_][a-zA-Z_0-9]*);

// Fixed: Removed SignPart from DecimalIntegerLiteral
fragment DecimalIntegerLiteral
  : DecimalNumeral
  ;

fragment HexIntegerLiteral
  : '0' [xX] HexDigits
  ;

fragment OctalIntegerLiteral
  : '0' OctalDigits
  ;

fragment BinaryIntegerLiteral
  : '0' [bB] BinaryDigits
  ;

fragment DecimalNumeral
  : '0'
  | NonZeroDigit DecimalDigits?
  ;

fragment DecimalDigits
  : DecimalDigit+
  ;

fragment DecimalDigit
  : [0-9]
  ;

fragment NonZeroDigit
  : [1-9]
  ;

fragment HexDigits
  : HexDigit+
  ;

fragment HexDigit
  : [0-9a-fA-F]
  ;

fragment OctalDigits
  : OctalDigit+
  ;

fragment OctalDigit
  : [0-7]
  ;

fragment BinaryDigits
  : BinaryDigit+
  ;

fragment BinaryDigit
  : [01]
  ;

fragment ExponentPart
  : [eE] [+-]? DecimalDigits
  ;

// Comments and whitespace
COMMENT
  : '//' ~[\r\n]* -> skip
  ;

BLOCK_COMMENT
  : '/*' .*? '*/' -> skip
  ;

WHITESPACE
  : [ \t\r\n\u000C\u00A0\u2028\u2029]+ -> skip
  ;
