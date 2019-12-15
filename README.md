# preStop に sleep 10 するだけで Deployment は、中断なく動くか

## 検証したいこと

- DeploymentをUpdateした場合、処理中のリクエストがあったとしても、SIGTERMが送られてPodは終了してしまう。その場合、クライアントにはサーバから切断されたように映る。
- これを、lifecycle.preStopに`/bin/sleep 10`を設定すると以下が期待できる
  - preStopが実行されたと同時に、LoadBalancerからの新しいリクエストを受け付けなくなる
  - 10秒の待機時間の間に、処理中のリクエストを処理する

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: graceful-deployment
spec:
  replicas: 8
  template:
    spec:
      containers:
      - name: graceful
        lifecycle:
          preStop:
            exec:
              command:
                - "/bin/sleep"
                - "10"
```

## 検証環境

- Google Kubernetes Engine を用いる
- Webアプリケーションを構築する。
  - /hello にGETリクエストを送ると、3秒後に応答する
- クライアントを構築する。
  - locust を用いて、100クライアントからの同時アクセスを行う
- Webアプリケーションは、NEGを適用したIngressを用いる
- WebアプリケーションのPodの数（replicas）は8

## 検証1 なにもしないと、Deploymentで処理中のリクエストは切断されるかどうか

- failer: 20 `HTTPError('502 Server Error: Bad Gateway for url: http://35.227.204.98/hello')` が発生

## 検証2

- failer: 0

## 結論

雑に lifecycle.preStop に `sleep 10` を指定するだけでも有用
