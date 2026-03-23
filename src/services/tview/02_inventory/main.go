package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/rivo/tview"
)

// Item - representação de uma entidade de produto no estoque
type Item struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

var (
	inventory     = []Item{}
	inventoryFile = "inventory.json"
)

// Carregamento do inventário a partir de um arquivo JSON persistido em disco
func loadInventory() {
	if _, err := os.Stat(inventoryFile); err == nil {
		data, err := os.ReadFile(inventoryFile)
		if err != nil {
			log.Fatal("Erro ao ler o arquivo de inventário:", err)
		}
		json.Unmarshal(data, &inventory)
	}
}

// Persistencia do estado atual do inventário em formato JSON
func saveInventory() {
	data, err := json.MarshalIndent(inventory, "", "  ")
	if err != nil {
		log.Fatal("Erro ao salvar o inventário:", err)
	}
	os.WriteFile(inventoryFile, data, 0644)
}

// Remoção de um item do inventário com base no índice informado
func deleteItem(index int) {
	if index < 0 || index >= len(inventory) {
		fmt.Println("Índice de item inválido.")
		return
	}

	inventory = append(inventory[:index], inventory[index+1:]...)
	saveInventory()
}

// Função principal responsável pela inicialização da aplicação TUI
func main() {
	app := tview.NewApplication()

	loadInventory()

	inventoryList := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	inventoryList.SetBorder(true).SetTitle("Itens do Inventário")

	// Atualização da visualização do inventário na interface
	refreshInventory := func() {
		inventoryList.Clear()

		if len(inventory) == 0 {
			fmt.Fprintln(inventoryList, "Nenhum item no inventário.")
		} else {
			for i, item := range inventory {
				fmt.Fprintf(inventoryList, "[%d] %s (Estoque: %d)\n", i+1, item.Name, item.Stock)
			}
		}
	}

	itemNameInput := tview.NewInputField().SetLabel("Nome do Item: ")
	itemStockInput := tview.NewInputField().SetLabel("Quantidade: ")
	itemIDInput := tview.NewInputField().SetLabel("ID do item para exclusão: ")

	form := tview.NewForm().
		AddFormItem(itemNameInput).
		AddFormItem(itemStockInput).
		AddFormItem(itemIDInput).
		AddButton("Adicionar Item", func() {
			name := itemNameInput.GetText()
			stock := itemStockInput.GetText()

			if name != "" && stock != "" {
				quantity, err := strconv.Atoi(stock)
				if err != nil {
					fmt.Fprintln(inventoryList, "Valor de estoque inválido.")
					return
				}

				inventory = append(inventory, Item{Name: name, Stock: quantity})
				saveInventory()
				refreshInventory()

				itemNameInput.SetText("")
				itemStockInput.SetText("")
			}
		}).
		AddButton("Remover Item", func() {
			idStr := itemIDInput.GetText()

			if idStr == "" {
				fmt.Fprintln(inventoryList, "Informe o ID do item para exclusão.")
				return
			}

			id, err := strconv.Atoi(idStr)
			if err != nil || id < 1 || id > len(inventory) {
				fmt.Fprintln(inventoryList, "ID do item inválido.")
				return
			}

			deleteItem(id - 1)
			fmt.Fprintf(inventoryList, "Item [%d] removido com sucesso.\n", id)
			refreshInventory()
			itemIDInput.SetText("")
		}).
		AddButton("Sair", func() {
			app.Stop()
		})

	form.SetBorder(true).SetTitle("Gerenciamento de Inventário").SetTitleAlign(tview.AlignLeft)

	flex := tview.NewFlex().
		AddItem(inventoryList, 0, 1, false).
		AddItem(form, 0, 1, true)

	refreshInventory()

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
