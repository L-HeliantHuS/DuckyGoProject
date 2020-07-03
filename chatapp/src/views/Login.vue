<template>
	<div class="Login">
		<el-container>
			<el-form :model="form" label-width="200px" ref="form">
				<el-row>
					<el-col :span="24">
						<el-form-item label="用户名">
							<el-input v-model="form.user_name"></el-input>
						</el-form-item>
					</el-col>
				</el-row>
				
				<el-row>
					<el-col :span="24">
						<el-form-item label="密码">
							<el-input type="password" v-model="form.password"></el-input>
						</el-form-item>
					</el-col>
				</el-row>
				
				<el-row>
					<el-col :span="12">
						<el-form-item>
							<el-button @click="submitLogin" type="primary" :loading="loading">提交</el-button>
						</el-form-item>
					</el-col>
				</el-row>
			</el-form>
		</el-container>
	</div>
</template>

<script>
  import {userLogin} from "../api/main";

  export default {
    name: "Login",
    data: () => ({
      form: {
        user_name: "",
        password: ""
      },
	  loading: false
    }),
    methods: {
      submitLogin() {
        this.loading = true;
        userLogin(this.form).then(
            response => {
              if (response.code === 0) {
                // 将获取的JWT Token放入LocalStorage
                localStorage.setItem("token", response.data.access_token);
                // 跳转到管理页面
                this.$router.push({name: "home"});
                // 并刷新网页
                window.location.reload();
              } else {
                // 显示提示 (用户名输入错误 或者其他的)
                this.$notify.warning({
                  message: response.msg
                });

				this.loading = false;
              }
            }
        )
      }
    },
	
  }
</script>

<style scoped>

</style>
