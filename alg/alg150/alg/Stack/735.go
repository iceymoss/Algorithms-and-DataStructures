package Stack

func asteroidCollision(asteroids []int) []int {
	//使用栈
	if len(asteroids) == 0 {
		return []int{}
	}
	sta := make([]int, 0)
	for i := 0; i < len(asteroids); i++ {
		//使用alive来标记行星是否存活
		alive := true
		if asteroids[i] >= 0 {
			sta = append(sta, asteroids[i])
		} else {
			for alive && len(sta) > 0 && sta[len(sta)-1] > 0 {
				alive = sta[len(sta)-1] < -asteroids[i]
				if sta[len(sta)-1] <= -asteroids[i] {
					sta = sta[:len(sta)-1]
				}
			}
			//当数组为向左运动
			if alive {
				sta = append(sta, asteroids[i])
			}
		}
	}
	return sta
}
