package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	ID       int
	Amount   float64
	Category string
	Date     time.Time
	Type     string
}

type BudgetTracker struct {
	transactions []Transaction
	nextID       int
}

type FinancialRecord interface {
	GetAmount() float64
	GetType() string
}

func (t Transaction) GetAmount() float64 {
	return t.Amount
}

func (t Transaction) GetType() string {
	return t.Type
}

func (bt *BudgetTracker) AddTransaction(amount float64, category, tType string) {
	newTransaction := Transaction{
		ID:       bt.nextID,
		Amount:   amount,
		Category: category,
		Date:     time.Now(),
		Type:     tType,
	}
	bt.transactions = append(bt.transactions, newTransaction)
	bt.nextID++
}

func (bt BudgetTracker) DisplayTransactions() {
	if len(bt.transactions) == 0 {
		fmt.Println("Nenhuma transação cadastrada.")
		return
	}

	fmt.Println("ID\tValor\tCategoria\tData\tTipo")
	for _, transaction := range bt.transactions {
		fmt.Printf("%d\t%.2f\t%s\t%s\t%s\n",
			transaction.ID,
			transaction.Amount,
			transaction.Category,
			transaction.Date.Format("02/01/2006 15:04"),
			transaction.Type,
		)
	}
}

func (bt BudgetTracker) CalculateTotal(tType string) float64 {
	var total float64
	for _, transaction := range bt.transactions {
		if transaction.Type == tType {
			total += transaction.Amount
		}
	}
	return total
}

func (bt BudgetTracker) SaveToCSV(filename string) error {
	if len(bt.transactions) == 0 {
		fmt.Println("Nenhuma transação para salvar.")
		return nil
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"ID", "Valor", "Categoria", "Data", "Tipo"}); err != nil {
		return err
	}

	for _, t := range bt.transactions {
		record := []string{
			strconv.Itoa(t.ID),
			strconv.FormatFloat(t.Amount, 'f', 2, 64),
			t.Category,
			t.Date.Format("02/01/2006 15:04"),
			t.Type,
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	fmt.Println("Transações salvas em", filename)
	return nil
}

func main() {
	bt := BudgetTracker{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Controle Financeiro ===")
		fmt.Println("1. Adicionar transação")
		fmt.Println("2. Listar transações")
		fmt.Println("3. Total de receitas")
		fmt.Println("4. Total de despesas")
		fmt.Println("5. Salvar em CSV")
		fmt.Println("6. Sair")
		fmt.Println("Escolha uma opção:")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Entrada inválida!")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Digite o valor: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			amount, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Valor inválido!")
				continue
			}

			fmt.Print("Digite a categoria (ex: alimentação, transporte, salário): ")
			category, _ := reader.ReadString('\n')
			category = strings.TrimSpace(category)

			if category == "" {
				fmt.Println("Categoria inválida!")
				continue
			}

			fmt.Print("Digite o tipo (receita/despesa): ")
			tType, _ := reader.ReadString('\n')
			tType = strings.TrimSpace(strings.ToLower(tType))

			if tType != "receita" && tType != "despesa" {
				fmt.Println("Tipo inválido!")
				continue
			}

			bt.AddTransaction(amount, category, tType)

		case 2:
			bt.DisplayTransactions()

		case 3:
			total := bt.CalculateTotal("receita")
			fmt.Printf("Total de receitas: R$ %.2f\n", total)

		case 4:
			total := bt.CalculateTotal("despesa")
			fmt.Printf("Total de despesas: R$ %.2f\n", total)

		case 5:
			fmt.Print("Digite o nome do arquivo: ")
			filename, _ := reader.ReadString('\n')
			filename = strings.TrimSpace(filename)

			err := bt.SaveToCSV(filename)
			if err != nil {
				fmt.Println("Erro ao salvar:", err)
			}

		case 6:
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida!")
		}
	}
}
