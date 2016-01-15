%{
package main


type Expression interface{}
type Token struct {
  token   int
  literal string
}
type Statement interface{}

type NumExpr struct {
  literal string
}
type StringExpr struct {
  literal string
}
type BinOpExpr struct {
  left     Expression
  operator rune
  right    Expression
}
type EchoStatement struct {
  expr Expression
}
%}

%union{
    token Token
    expr  Expression
    statement Statement
}

%type<expr> program
%type<expr> expr
%type<statement> statement
%token<token> NUMBER STRING ECHO

%left '+'

%%

program
    : expr ';'
    | statement ';'
    {
        $$ = $1
        yylex.(*Lexer).result = $$
    }
statement
    : ECHO expr
    {
      $$ = EchoStatement{expr: $2}
    }
expr
    : NUMBER
    {
        $$ = NumExpr{literal: $1.literal}
    }
    | STRING
    {
        $$ = StringExpr{literal: $1.literal}
    }
    | expr '+' expr
    {
        $$ = BinOpExpr{left: $1, operator: '+', right: $3}
    }
%%
