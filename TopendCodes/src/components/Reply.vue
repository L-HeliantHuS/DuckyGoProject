<template>
	<div class="reply">
		<CreateReply :displaySync="dp" :closeHandler="closeHandlerFunc" :oid=replyConfig.oid :rid=replyConfig.rid
					 :root=replyConfig.root :userName="replyConfig.userName"></CreateReply>
		
		
		<div v-for="(item, index) in replys" :key="item.id">
			<br>
			{{ index + 1 }} 楼 | 共有{{ item.count }}条回复
			<div class="Root">
				{{ item.person.user_name }} : {{ item.content }}
				
				<button v-on:click="replyInfoInsert(id, item.id, item.id, item.person.user_name)">回复</button>
			</div>
			
			
			<div class="Child" v-for="child in item.members">
				<div v-if="child.rid === item.id">
					{{ child.person.user_name }}: {{ child.content }}
					<button v-on:click="replyInfoInsert(id, child.id, item.id, child.person.user_name)">回复</button>
				
				</div>
				
				<div v-else>
					{{ child.person.user_name }} 回复 {{ getUserInfo(index, child.rid).person.user_name }} : {{
					child.content }}
					<button v-on:click="replyInfoInsert(id, child.id, item.id, child.person.user_name)">回复</button>
				
				</div>
			</div>
			<hr>
		</div>
	</div>

</template>

<script>
  import * as API from "@/api/api"
  import CreateReply from "@/components/CreateReply";

  export default {
    name: "Reply",
    components: {CreateReply},
    props: {
      id: 0,
    },

    data() {
      return {
        replys: [],
        userList: {},
        dp: false,

        replyConfig: {
          oid: 0,
          rid: 0,
          root: 0,
          userName: "",
        }
      }
    },

    created() {
      this.getReply();
    },

    methods: {

      getReply() {
        API.replyGet(this.id)
            .then(response => {
              if (response.code === 0) {
                this.replys = response.data;
              } else {
                this.$message.error(`获取数据异常, ${response.msg}`)
              }
            })
      },

      getUserInfo(index, id) {
        return this.replys[index].members.find(item => {
          return item.id === id
        })
      },


      replyInfoInsert(oid, rid, root, userName) {
        this.replyConfig.oid = Number(oid);
        this.replyConfig.rid = rid;
        this.replyConfig.root = root;
        this.replyConfig.userName = userName;


        this.dp = true;
      },

      closeHandlerFunc() {
        this.dp = false;
      }
    }
  }
</script>

<style scoped>
	.Root {
		width: 100%;
		height: 50px;
		text-align: left;
	}
	
	
	.Child {
		width: 70%;
		margin-left: 30px;
		text-align: left;
	}
</style>
