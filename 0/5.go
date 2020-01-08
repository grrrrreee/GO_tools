package main

// type struct를 통해서 거래의 형태를 정의함
type Transaction struct {
	amount int
	ID     []byte
}

// Transaction을 tx로 따와서 해당 tx의 amount 가 0 보다 큰지에 대해서 확인
func isValid(tx *Transaction) {
	return tx.amount > 0
}
