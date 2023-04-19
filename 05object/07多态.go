package main

import "fmt"

/*
多态：一个父类接口可以由多个子类实现
*/

type Rider struct {
}
type Archer struct {
}
type Master struct {
}

// Attack 骑兵
func (r *Rider) Attack() (bloodloss int) {
	fmt.Println("铁蹄滚滚，尸横遍野")
	return 100
}
func (r *Rider) Defend() {
	fmt.Println("骑兵防守")
}

// 弓箭手
func (r *Archer) Attack() (bloodloss int) {
	fmt.Println("万箭齐发")
	return 100
}
func (r *Archer) Defend() {
	fmt.Println("弓箭手防守")
}

// 法师
func (r *Master) Attack() (bloodloss int) {
	fmt.Println("师法天地")
	return 100
}
func (r *Master) Defend() {
	fmt.Println("法师防守")
}

func main071() {
	fighters := make([]Fighter, 0)
	fighters = append(fighters, &Rider{})
	fighters = append(fighters, &Archer{})
	fighters = append(fighters, &Master{})
	/*
		用户发将令
		000-全体防守
		999-全体进攻
		090-弓箭手进攻，其他防守
	*/
	var cmd string
	fmt.Println("请发将令")
	fmt.Scan(&cmd)
	switch cmd {
	case "999":
		for _, fighter := range fighters {
			fighter.Attack()
		}
	case "000":
		for _, fighter := range fighters {
			fighter.Defend()
		}
	default:
		for _, f := range fighters {
			/*
				// 第二种类型断言
				// 调度骑兵
				if rider, ok := f.(*Rider); ok {
					if cmd[0] == '9' {
						rider.Attack()
					} else {
						rider.Defend()
					}
				}
				// 调度弓箭手
				if archer, ok := f.(*Archer); ok {
					if cmd[1] == '9' {
						archer.Attack()
					} else {
						archer.Defend()
					}
				}
				// 调度法师
				if master, ok := f.(*Master); ok {
					if cmd[2] == '9' {
						master.Attack()
					} else {
						master.Defend()
					}
				}

			*/
			// switch类型断言
			switch f.(type) {
			case *Archer:
				if cmd[1] == '9' {
					f.Attack()
				} else {
					f.Defend()
				}
			case *Rider:
				if cmd[0] == '9' {
					f.Attack()
				} else {
					f.Defend()
				}
			case *Master:
				if cmd[2] == '9' {
					f.Attack()
				} else {
					f.Defend()
				}
			}
		}
	}

}
