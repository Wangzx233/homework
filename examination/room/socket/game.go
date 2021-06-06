package socket

const size = 30
func cp(first,second int,nowPlayer string,roomID string) string {
	if Checkerboard[roomID][first][second]==""{
		i := Checkerboard[roomID]
		i[first][second]=nowPlayer
		Checkerboard[roomID]=i



		if isWin(first,second,nowPlayer,i){

			Mng.GameOver(Message{
				Type:    "game_over",
				User:    nowPlayer,
				Content: "",
				RoomID:  roomID,
			})
			return "0"
		}

		if nowPlayer=="1" {
			return "2"
		}else {
			return "1"
		}
	}else {
		return nowPlayer
	}
}
func isWin(first,second int,nowPlayer string,cb [size][size]string) bool {
	for i:=1;i<=5; {
		//向左
		for (cb[first][second]==nowPlayer)&&(first-i>=0) {
			if cb[first-i][second]==nowPlayer {
				i++
			}else {
				break
			}
		}
		//向右
		for (cb[first][second]==nowPlayer)&&(first+i<size) {
			if cb[first+i][second]==nowPlayer {
				i++
			}else {
				break
			}
		}
		if i>=5 {
			break
		}
		i=1

		//左上
		for (cb[first][second]==nowPlayer)&&(first-i>=0)&&(second-i>=0) {
			if cb[first-i][second-i]==nowPlayer {
				i++
			}else {
				break
			}
		}
		//右下
		for (cb[first][second]==nowPlayer)&&(first+i<size)&&(second<size) {
			if cb[first+i][second+i]==nowPlayer {
				i++
			}else {
				break
			}
		}
		if i>=5 {
			break
		}
		i=1

		//上
		for (cb[first][second]==nowPlayer)&&(second+i<size) {
			if cb[first][second+i]==nowPlayer {
				i++
			}else {
				break
			}
		}
		//下
		for (cb[first][second]==nowPlayer)&&(second-i>=0) {
			if cb[first][second-i]==nowPlayer {
				i++
			}else {
				break
			}
		}
		if i>=5 {
			break
		}
		i=1

		//右上
		for (cb[first][second]==nowPlayer)&&(first+i<size)&&(second-i>=0) {
			if cb[first+i][second-i]==nowPlayer {
				i++
			}else {
				break
			}
		}
		//左下
		for (cb[first][second]==nowPlayer)&&(first-i>=0)&&(second+i<size) {
			if cb[first-i][second+i]==nowPlayer {
				i++
			}else {
				break
			}
		}
		if i>=5 {
			break
		}
		i=1


		return false
	}
	return true
}