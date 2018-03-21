## Moment

date utils

## API

three way get `Mymoment` struct

```
import "github.com/huyinghuan/moment"
monent.Now()
monent.New(t *time.Time)
moment.NewFromUinx(unixTime int64)
```


### Mymoment struct

```
type Mymonent struct {
	Unix       int64
	Now        time.Time
}
```

#### func

