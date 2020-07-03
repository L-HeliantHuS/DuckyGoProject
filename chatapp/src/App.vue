<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
    </div>
    
    <div v-if="user.nickname !== ''">
      {{ user.nickname }} <el-button type="danger" plain @click="LogoutUser">Logout</el-button>
      <br>
    </div>
    
    <router-view/>
  </div>
</template>

<script>
  import {userLogout, userMe} from "./api/main";

  export default {
    data: () => ({
      user: {
        nickname: "",
      }
    }),
    created() {
      userMe()
      .then(response => {
        if (response.code === 0) {
          console.log(response.data);
          this.user.nickname = response.data.user.nickname
        }
      })
    },
    methods: {
      LogoutUser() {
        userLogout()
        .then(response => {
          this.$message.success("注销成功！");
          localStorage.removeItem("token");
          window.location.reload();
        })
      },
    }
  }
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
