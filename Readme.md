# Golang handler of OS signals

The code helps you organise waiting and handling OS signals like for example: SIGINT, SIGTERM, SIGKILL

## Create new handler

```go
import sig "github.com/leprosus/golang-sig"

sig.WaitSig(func(){
	fmt.Println("Buy-buy World!")
}, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
```

NB: default handling signals are SIGINT, SIGTERM, SIGKILL.

It means you don't need to set them in the function if you need to work with them only.

You have to set them in a case of an extended signal list.

The code above will wait getting one of SIGINT, SIGTERM, SIGKILL from OS. After getting a signal callback will be executed.