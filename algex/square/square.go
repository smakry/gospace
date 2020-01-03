package square

// purpose: 判定给定的数n是某个数的平方
// mothods: 1.从1到n判定是不是这个数开方
// 			2.在方法1的基础上二分方式，减少循环次数
//			3.使用数学归纳法： 
//				某个数的平方:	1^2,2^2,3^2,4^2,5^2,...,n
//							1  ,4  ,9  ,16 ,25 ,...,n
//					差值:	1  ,3  ,5  ,7  ,9  ,... //公差为2的等差数列
// 					结论：	n从等差数列1开始减，若刚好为0表示是某个数的平方，若小于0则表示不是

func SquareIsN(n int) (bool, int) {
	if n <= 0 {
		return false, 0
	}

	for i, m := 1, 1; ; {
		n -= i
		if n == 0 {
			return true, m
		} else if n < 0 {
			return false, 0
		} else {
			i += 2
			m++
		}
	}
}


// purpose: 判断一个数m是不是2的n次方
// mothods: 1.通过构造计算2的n次方判断是不是等于数m
// 			2.如果一个数是2的n次方，那么对应的二进制数只有1位是1其余为0，又m-1是m二进制1之后所有位为1与m相反

func IsOf2Square(m int) bool {
	if m <= 0 {
		return false
	}

	return m & (m - 1) == 0
}






