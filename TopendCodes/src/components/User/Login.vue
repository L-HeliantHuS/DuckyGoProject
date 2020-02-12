<template>
	<div class="login">
	
	</div>
</template>

<script>
  import * as API from "../../api/api";

  export default {
    name: "Login",
    data() {
      return {
        userdata: {
          user_name: "helianthus",
          password: "testpassword."
        },
      }
    },


    created() {
      API.userLogin(this.userdata)
          .then(response => {
            if (response.code === 0) {
              this.$message.success("登录成功")
            } else {
              this.$message.error("登录失败")
            }
          })

    }


  }
</script>

<style scoped>

</style>
