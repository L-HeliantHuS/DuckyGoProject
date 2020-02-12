<template>
	<div class="createReply">
		
		<el-dialog
				title="提示"
				width="30%"
				:visible.sync="displaySync"
				:before-close="closeHandler"
		>
			<span>这是一段信息</span>
			{{ oid }}
			{{ rid }}
			{{ root }}
			
			{{ content }}
			<el-input :placeholder="'回复给: ' + userName" type="textarea" v-model="content">
			
			</el-input>
			<br>
			<el-button type="success" v-on:click="CreatedReply">提交</el-button>
		</el-dialog>
	
	
	</div>
</template>

<script>
  import * as API from "@/api/api";

  export default {
    name: "CreateReply",

    props: {
      displaySync: Boolean,
      closeHandler: Function,

      oid: 0,
      rid: Number,
      root: Number,
      userName: String
    },

    data() {
      return {
        content: "",
      }
    },

    methods: {
      CreatedReply() {
        let form = {
          oid: this.oid,
          rid: this.rid,
          root: this.root,
          content: this.content
        };

        API.replyCreate(form)
            .then(response => {
              if (response.code === 0) {
                this.$message.success("评论成功");
				window.location.reload();
              } else {
                this.$message.error("评论失败")
              }
            })
      }
    }

  }
</script>

<style scoped>

</style>
