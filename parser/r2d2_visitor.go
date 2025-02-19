// Code generated from r2d2.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // r2d2

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by r2d2Parser.
type r2d2Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by r2d2Parser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by r2d2Parser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by r2d2Parser#globalDeclaration.
	VisitGlobalDeclaration(ctx *GlobalDeclarationContext) interface{}

	// Visit a parse tree produced by r2d2Parser#importDeclaration.
	VisitImportDeclaration(ctx *ImportDeclarationContext) interface{}

	// Visit a parse tree produced by r2d2Parser#interfaceDeclaration.
	VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{}

	// Visit a parse tree produced by r2d2Parser#moduleDeclaration.
	VisitModuleDeclaration(ctx *ModuleDeclarationContext) interface{}

	// Visit a parse tree produced by r2d2Parser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by r2d2Parser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by r2d2Parser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by r2d2Parser#typeExpression.
	VisitTypeExpression(ctx *TypeExpressionContext) interface{}

	// Visit a parse tree produced by r2d2Parser#arrayDimensions.
	VisitArrayDimensions(ctx *ArrayDimensionsContext) interface{}

	// Visit a parse tree produced by r2d2Parser#baseType.
	VisitBaseType(ctx *BaseTypeContext) interface{}

	// Visit a parse tree produced by r2d2Parser#genericType.
	VisitGenericType(ctx *GenericTypeContext) interface{}

	// Visit a parse tree produced by r2d2Parser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by r2d2Parser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by r2d2Parser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by r2d2Parser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by r2d2Parser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by r2d2Parser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by r2d2Parser#assignmentDeclaration.
	VisitAssignmentDeclaration(ctx *AssignmentDeclarationContext) interface{}

	// Visit a parse tree produced by r2d2Parser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by r2d2Parser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by r2d2Parser#simpleFor.
	VisitSimpleFor(ctx *SimpleForContext) interface{}

	// Visit a parse tree produced by r2d2Parser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by r2d2Parser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by r2d2Parser#loopControl.
	VisitLoopControl(ctx *LoopControlContext) interface{}

	// Visit a parse tree produced by r2d2Parser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by r2d2Parser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by r2d2Parser#logicalExpression.
	VisitLogicalExpression(ctx *LogicalExpressionContext) interface{}

	// Visit a parse tree produced by r2d2Parser#comparisonExpression.
	VisitComparisonExpression(ctx *ComparisonExpressionContext) interface{}

	// Visit a parse tree produced by r2d2Parser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by r2d2Parser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by r2d2Parser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by r2d2Parser#memberExpression.
	VisitMemberExpression(ctx *MemberExpressionContext) interface{}

	// Visit a parse tree produced by r2d2Parser#memberPart.
	VisitMemberPart(ctx *MemberPartContext) interface{}

	// Visit a parse tree produced by r2d2Parser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by r2d2Parser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by r2d2Parser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by r2d2Parser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by r2d2Parser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by r2d2Parser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by r2d2Parser#switchCase.
	VisitSwitchCase(ctx *SwitchCaseContext) interface{}

	// Visit a parse tree produced by r2d2Parser#defaultCase.
	VisitDefaultCase(ctx *DefaultCaseContext) interface{}
}
