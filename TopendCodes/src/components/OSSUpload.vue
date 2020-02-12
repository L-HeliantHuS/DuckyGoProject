<template>
	<div class="ossupload">
		<el-upload
				class="avatar-uploader"
				ref="upload"
				action=""
				:before-upload="fileChange"
				:show-file-list="false"
				:http-request="uploadToOSS">
			<img v-if="resultURL" :src="resultURL" class="avatar" crossOrigin="anonymous">
			<i v-else class="el-icon-plus avatar-uploader-icon"/>
			<div class="el-upload__tip">只能上传PNG文件, 且不超过500kb</div>
		</el-upload>
	</div>
</template>

<script>
  import * as API from '../api/api'

  export default {
    name: "OSSUpload",
    data() {
      return {

        form: {
          filename: ""
        },

        uploadURL: "",
        resultURL: ""
      }
    },
    methods: {

      fileChange(file) {
        const isPNG = file.type === 'image/png';

        if (!isPNG) {
          this.$message.error("上传文件必须是PNG格式")
        }


        this.form.filename = file.name;

        return isPNG
      },

      submitUpload() {
        this.$refs.upload.submit();
      },

      uploadToOSS(options) {
        // 获取OSS的key
        API.ossGetKey(this.form)
            .then(response => {
                  if (response.code !== 0) {
                    this.$message.error(
                        response.msg
                    )
                  } else {
                    const xmlHttpRequest = new XMLHttpRequest();
                    xmlHttpRequest.open("PUT", response.data.put, true);
                    xmlHttpRequest.setRequestHeader("HeliantHuS", "value");
                    xmlHttpRequest.send(options.file);
                    xmlHttpRequest.onload = () => {
                      this.resultURL = response.data.get;
                      this.from.filename = response.data.upload;
                    };
                  }
                }
            )


      }
    }

  }
</script>

<style>
	.avatar-uploader .el-upload {
		border: 1px dashed #d9d9d9;
		border-radius: 6px;
		cursor: pointer;
		position: relative;
		overflow: hidden;
	}
	
	.avatar-uploader .el-upload:hover {
		border-color: #409EFF;
	}
	
	.avatar-uploader-icon {
		font-size: 28px;
		color: #8c939d;
		width: 178px;
		height: 178px;
		text-align: center;
		margin: 30px;
		
	}
	
	.avatar {
		width: 178px;
		height: 178px;
		display: block;
	}
</style>
