# Golang Questions

可以借鑑抄襲或借助AI等工具，但是必須理解並解釋答案。

用例在main.go

除定義的TODO以外，可以根據自己對用法的理解去拓展

## Question 1 定義error

請完成`errorx.go`文件，使其能夠正確運行。

這是用於全局錯誤處理的庫，主要特點
- 實現stack的捕獲以方便頂層處理log
- 帶有code和type，方便進行業務邏輯判斷

## Question 2 定義config

請完成`config.go`文件，使其能夠正確運行。

這是用於定義配置的庫，主要特點
- 通過泛型約束配置檔案的格式
- 方便進行配置的crud操作

`local`或`etcd`實現任意一個即可