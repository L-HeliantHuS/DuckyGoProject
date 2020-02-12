<template>
	<div class="home">
		<Login></Login>
 		<reply id=1></reply>
	</div>
</template>

<script>
  // @ is an alias to /src

  import Reply from "@/components/Reply";
  import CreateReply from "@/components/CreateReply";
  import Login from "@/components/User/Login";

  export default {
    name: 'home',
    components: {
      Reply, CreateReply, Login
    }
  }
</script>
