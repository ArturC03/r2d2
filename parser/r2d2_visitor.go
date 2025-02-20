// Code generated from R2D2.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // R2D2

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by R2D2Parser.
type R2D2Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by R2D2Parser#program.
	VisitProgram(ctx *ProgramContext) any

	// Visit a parse tree produced by R2D2Parser#declaration.
	VisitDeclaration(ctx *DeclarationContext) any

	// Visit a parse tree produced by R2D2Parser#globalDeclaration.
	VisitGlobalDeclaration(ctx *GlobalDeclarationContext) any

	// Visit a parse tree produced by R2D2Parser#importDeclaration.
	VisitImportDeclaration(ctx *ImportDeclarationContext) any

	// Visit a parse tree produced by R2D2Parser#interfaceDeclaration.
	VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) any

	// Visit a parse tree produced by R2D2Parser#moduleDeclaration.
	VisitModuleDeclaration(ctx *ModuleDeclarationContext) any

	// Visit a parse tree produced by R2D2Parser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) any

	// Visit a parse tree produced by R2D2Parser#parameterList.
	VisitParameterList(ctx *ParameterListContext) any

	// Visit a parse tree produced by R2D2Parser#parameter.
	VisitParameter(ctx *ParameterContext) any

	// Visit a parse tree produced by R2D2Parser#typeExpression.
	VisitTypeExpression(ctx *TypeExpressionContext) any

	// Visit a parse tree produced by R2D2Parser#arrayDimensions.
	VisitArrayDimensions(ctx *ArrayDimensionsContext) any

	// Visit a parse tree produced by R2D2Parser#baseType.
	VisitBaseType(ctx *BaseTypeContext) any

	// Visit a parse tree produced by R2D2Parser#genericType.
	VisitGenericType(ctx *GenericTypeContext) any

	// Visit a parse tree produced by R2D2Parser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) any

	// Visit a parse tree produced by R2D2Parser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) any

	// Visit a parse tree produced by R2D2Parser#statement.
	VisitStatement(ctx *StatementContext) any

	// Visit a parse tree produced by R2D2Parser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) any

	// Visit a parse tree produced by R2D2Parser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) any

	// Visit a parse tree produced by R2D2Parser#forStatement.
	VisitForStatement(ctx *ForStatementContext) any

	// Visit a parse tree produced by R2D2Parser#assignmentDeclaration.
	VisitAssignmentDeclaration(ctx *AssignmentDeclarationContext) any

	// Visit a parse tree produced by R2D2Parser#assignment.
	VisitAssignment(ctx *AssignmentContext) any

	// Visit a parse tree produced by R2D2Parser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) any

	// Visit a parse tree produced by R2D2Parser#simpleFor.
	VisitSimpleFor(ctx *SimpleForContext) any

	// Visit a parse tree produced by R2D2Parser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) any

	// Visit a parse tree produced by R2D2Parser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) any

	// Visit a parse tree produced by R2D2Parser#loopControl.
	VisitLoopControl(ctx *LoopControlContext) any

	// Visit a parse tree produced by R2D2Parser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) any

	// Visit a parse tree produced by R2D2Parser#expression.
	VisitExpression(ctx *ExpressionContext) any

	// Visit a parse tree produced by R2D2Parser#logicalExpression.
	VisitLogicalExpression(ctx *LogicalExpressionContext) any

	// Visit a parse tree produced by R2D2Parser#comparisonExpression.
	VisitComparisonExpression(ctx *ComparisonExpressionContext) any

	// Visit a parse tree produced by R2D2Parser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) any

	// Visit a parse tree produced by R2D2Parser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) any

	// Visit a parse tree produced by R2D2Parser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) any

	// Visit a parse tree produced by R2D2Parser#memberExpression.
	VisitMemberExpression(ctx *MemberExpressionContext) any

	// Visit a parse tree produced by R2D2Parser#memberPart.
	VisitMemberPart(ctx *MemberPartContext) any

	// Visit a parse tree produced by R2D2Parser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) any

	// Visit a parse tree produced by R2D2Parser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) any

	// Visit a parse tree produced by R2D2Parser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) any

	// Visit a parse tree produced by R2D2Parser#literal.
	VisitLiteral(ctx *LiteralContext) any

	// Visit a parse tree produced by R2D2Parser#block.
	VisitBlock(ctx *BlockContext) any

	// Visit a parse tree produced by R2D2Parser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) any

	// Visit a parse tree produced by R2D2Parser#switchCase.
	VisitSwitchCase(ctx *SwitchCaseContext) any

	// Visit a parse tree produced by R2D2Parser#defaultCase.
	VisitDefaultCase(ctx *DefaultCaseContext) any
}
