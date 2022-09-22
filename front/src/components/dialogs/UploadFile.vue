<template>
  <v-dialog v-model="showUploadDialog" persistent max-width="600px">
    <v-card>
      <v-card-title>
        <span class="text-h5">Upload file</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12">
              <v-file-input
                multiple
                label="Выберите файл для загрузки"
                ref="file"
                type="file"
                id="file"
                v-model="uploadFiles"
                @change="HandleFileChanged"
                value="file"
              >
              </v-file-input>
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
          color="blue darken-1"
          text
          @click="CloseUploadFileDialog"
        >
          Close
        </v-btn>
        <v-btn color="blue darken-1" text @click="HandleFileUpload">
          Upload
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "UploadFile",
  props: ["showUploadDialog", 'bucketName'],
  data: () => ({
    uploadFiles: [],
    fileLinks: [],
    uploadFilesProgress: [],
  }),
  methods: {
    CloseUploadFileDialog(){
      this.uploadFiles = undefined
      this.$emit('CloseUploadDialog')
    },
    HandleFileChanged() {
      this.$http
        .post("http://localhost:8080" + "/api/file/upload", {
          bucketName: this.bucketName,
          fileName: this.uploadFiles.map((el) => el.name),
        })
        .then((response) => {
          response.data.map((el) =>
            this.fileLinks.push({ fileName: el["fileName"], url: el["url"] })
          );
          console.log(this.fileLinks.map((x) => x.fileName));
        });
    },
    //TODO!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
    async HandleFileUpload() {
      await this.fileLinks.forEach((fileLink) => {
        let formdata = new FormData();
        this.uploadFiles.forEach((upFile) => {
          if (fileLink["fileName"] == upFile["name"]) {
            formdata.append("data", upFile);
          }
        });
        this.uploadFilesProgress.push({
          fileName: fileLink["fileName"],
          bucketName: this.bucketName,
          uploaded: 0,
        });
        this.$http
          .put(fileLink["url"], formdata, {
            uploadProgress: (upProgress) => {
              this.uploadFilesProgress.forEach((uploadFile)=>{
                if (uploadFile["fileName"]==fileLink["fileName"]){
                  uploadFile['uploaded'] = Math.round(
                    (upProgress.loaded * 100) / upProgress.total
                    )
                    console.log(uploadFile['fileName'], uploadFile["uploaded"])
                }
              })
            },
          })
          .then(() => {
            console.log("File", fileLink["fileName"], "is uploaded");
            this.uploadFilesProgress.splice(this.uploadFilesProgress.indexOf(fileLink["fileName"]),1)
            this.$root.$emit("CheckFilesUpdates")
          });
      });
      this.uploadDialog = false;
      this.uploadFiles = undefined;
      console.log("test", this.uploadFilesProgress)
      this.$emit("CloseUploadDialog");
    },
    //END TODO!!!!!!!!!!!!!!!!!!!!!!!
  },

};
</script>
