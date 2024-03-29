package js

import "github.com/dop251/goja/ast"

// JsStruct JavaScript 数据结构 可以将多个结构体组合在一起
type JsStruct struct {
	FunctionName []string // 函数名
	Parameter    []string // 入参名
	Variable     []string // 变量名
	Literal      []string // 字面量
}

// ParseJsAST 解析js的ast
func (j *JsStruct) ParseJsAST(prg *ast.Program) {

	for _, i := range prg.Body {
		j.BlockParsing(i)

	}

}

// BlockParsing 块解析 解析所有函数名,入参名 , 变量名 , 字面量
// [?] 后续需要更完整的解析 为用于构建平行镜像 创造更多的可能性
// 可以通过索引定位各个实例的位置 并且可以跟踪或者修改
func (j *JsStruct) BlockParsing(a any) {
	switch v := a.(type) {
	case *ast.FunctionDeclaration: // 函数声明
		j.FunctionName = append(j.FunctionName, string(v.Function.Name.Name))
		for _, p := range v.Function.ParameterList.List {
			identifier, ok := p.Target.(*ast.Identifier)
			if ok {
				j.Parameter = append(j.Parameter, string(identifier.Name))
			}
		}

	case *ast.VariableStatement: // 变量声明

		list := v.List
		for _, binding := range list {
			//binding.Target    // 新的语句
			j.BlockParsing(binding)
		}

	case *ast.ExpressionStatement: // 表达式声明
		j.BlockParsing(v.Expression)

	case *ast.BlockStatement: // 块声明
		for _, statement := range v.List {
			j.BlockParsing(statement)
		}

	case *ast.CallExpression: // 调用表达式

	case *ast.Identifier: // 标识符

	case *ast.DotExpression: // 点表达式

	case *ast.ObjectLiteral: // 对象字面量

	case *ast.FunctionLiteral: // 函数字面量
	case *ast.LexicalDeclaration: // 词法声明
	case *ast.ReturnStatement: // 返回声明

	case *ast.PropertyShort: // 属性短语
	case *ast.Binding: // 绑定
	case *ast.SuperExpression: // 超级表达式

	case *ast.TemplateElement: // 模板元素
	case *ast.SwitchStatement: // 开关声明

	case *ast.ForLoopInitializerExpression: // 循环初始化表达式
	case *ast.DebuggerStatement: // 调试声明
	case *ast.PropertyKeyed: // 属性键
	case *ast.ForLoopInitializerLexicalDecl: // 循环初始化词法声明
	case *ast.AwaitExpression: // 等待表达式
	case *ast.ForIntoVar: // 循环进入变量
	case *ast.ForDeclaration: // 循环声明
	case *ast.AssignExpression: // 赋值表达式
	case *ast.ClassLiteral: // 类字面量
	case *ast.BooleanLiteral: // 布尔字面量
	case *ast.SpreadElement: // 扩展元素
	case *ast.ForLoopInitializerVarDeclList: // 循环初始化变量声明列表
	case *ast.BinaryExpression: // 二进制表达式
	case *ast.ForOfStatement: // 循环的声明
	case *ast.NewExpression: // 新表达式
	case *ast.YieldExpression: // 产量表达式
	case *ast.ObjectPattern: // 对象模式
	case *ast.ExpressionBody: // 表达式体
	case *ast.NullLiteral: // 空字面量

	case *ast.IfStatement: // 如果声明
	case *ast.FieldDefinition: // 字段定义
	case *ast.BracketExpression: // 括号表达式
	case *ast.TryStatement: // 尝试声明
	case *ast.ArrayPattern: // 数组模式
	case *ast.PrivateIdentifier: // 私有标识符
	case *ast.CaseStatement: // 案例声明
	case *ast.ArrowFunctionLiteral: // 箭头函数字面量
	case *ast.CatchStatement: // 捕获声明
	case *ast.StringLiteral: // 字符串字面量
	case *ast.ClassStaticBlock: // 类静态块
	case *ast.BadExpression: // 错误表达式
	case *ast.PropertyKind: // 属性种类
	case *ast.OptionalChain: // 可选链
	case *ast.BadStatement: // 错误声明
	case *ast.WithStatement: // 与声明
	case *ast.MethodDefinition: // 方法定义
	case *ast.RegExpLiteral: // 正则表达式字面量
	case *ast.DoWhileStatement: // 做当声明
	case *ast.MetaProperty: // 元属性
	case *ast.SequenceExpression: // 序列表达式
	case *ast.ForInStatement: // 循环声明
	case *ast.ThisExpression: // 这个表达式
	case *ast.BranchStatement: // 分支声明
	case *ast.ParameterList: // 参数列表
	case *ast.ConditionalExpression: // 条件表达式
	case *ast.EmptyStatement: // 空声明
	case *ast.ThrowStatement: // 抛出声明
	case *ast.TemplateLiteral: // 模板字面量
	case *ast.ClassElement: // 类元素
	case *ast.Optional: // 可选

	case *ast.NumberLiteral: // 数字字面量
	case *ast.VariableDeclaration: // 变量声明
	case *ast.UnaryExpression: // 一元表达式
	case *ast.LabelledStatement: // 标记声明
	case *ast.ForIntoExpression: // 循环进入表达式
	case *ast.ForStatement: // 循环声明

	case *ast.WhileStatement: // 当声明

	case *ast.ClassDeclaration: // 类声明
	case *ast.ArrayLiteral: // 数组字面量
	case *ast.PrivateDotExpression: // 私有点表达式

	case *ast.Program: // 程序
	}

}

// FunctionDeclaration 函数声明
type FunctionDeclaration struct {
	FunctionName string   // 函数名
	Parameter    []string // 入参名
}

// Binding 绑定器 用于绑定变量
type Binding struct {
	Target      string // 目标
	Initializer string // 初始化
}
