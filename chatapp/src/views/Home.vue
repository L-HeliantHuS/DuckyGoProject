<template>
	<div class="home">
		<el-row>
			<el-input
				:autosize="{ minRows: 20, maxRows: 20}"
				placeholder="Listening..."
				style="resize:none;"
				type="textarea"
				v-model="textarea1">
			</el-input>
		</el-row>
		
		<el-row>
			<div style="margin-top: 15px;">
				<el-input class="input-with-select" placeholder="请输入内容" v-model="input">
					<el-button :loading="loading" @click="sendMessage" icon="el-icon-position"
							   slot="append"></el-button>
				</el-input>
			</div>
		</el-row>
	</div>
</template>

<script>

  import {userMe} from "../api/main";

  export default {
    name: 'Home',
    components: {},
    data: () => ({
      textarea1: "",
      input: "",
      user: {
        nickname: "",
        websock: null,
      },
      loading: false
    }),

    created() {
      // 检查是否登录
      this.checkLogin();

      this.initWebSocket();
    },
    methods: {
      checkLogin() {
        let token = localStorage.getItem("token");
        if (token == null) {
          this.$router.push("/user/login");
          this.$message.warning("你还没有登录, 请先登录.");
          return
        }

        userMe()
            .then(response => {
              this.user = response.data
            });
      },
      initWebSocket() {
        var loc = window.location, new_uri;
        if (loc.protocol === "https:") {
          new_uri = "wss:";
        } else {
          new_uri = "ws:";
        }
        new_uri += "//" + loc.host;
        new_uri += loc.pathname + "api/v2/ws?token=" + localStorage.getItem("token");
        
        const sockuri = new_uri;
        this.websock = new WebSocket(sockuri);
        this.websock.onmessage = this.websocketonmessage;
        this.websock.onopen = this.websocketonopen;
        this.websock.onerror = this.websocketonerror;
        this.websock.onclose = this.websocketclose;
      },

      websocketonopen() { //连接建立之后执行send方法发送数据
        let actions = "ping";
        this.websocketsend(actions);
      },
      websocketonerror() {//连接建立失败重连
        this.initWebSocket();
      },
      websocketonmessage(e) { //数据接收
        const redata = e.data;
        if (redata === "pong") {
          this.$message.success("WebSocket服务正常.")
        } else {
          let data = JSON.parse(redata);
          // 服务不正常的情况下提示
          if ( data.code !== 0 ) {
            this.$message.warning(data.msg)
          } else {
            this.textarea1 += data.timestamp + "\t" + data.data + "\n"
          }
        }
      },
      websocketsend(Data) {//数据发送
        this.websock.send(Data);
      },
      websocketclose(e) {  //关闭
        console.log('断开连接', e);
      },

      sendMessage() {
        this.loading = true;
        let data = {"message": this.input};
        this.websocketsend(JSON.stringify(data));
        this.input = "";
        
        this.loading = false;

      }

    }

  }
</script>
