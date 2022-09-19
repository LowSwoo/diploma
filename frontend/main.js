const server = {
  scheme: "http",
  host: "localhost",
  port: "8080",
};
const server_url = server.scheme + "://" + server.host + ":" + server.port;
new Vue({
  el: "#app",
  data: {
    bucket: {
      bucketName: "",
      bucketRegion: "",
      objectLocking: false,
    },

    bucketName: undefined,
    bucketList: [],
    createDialog: false,
    uploadDialog: false,
    uploadSnackbar: false,
    uploadFiles: [],
    files: [],
    file: null,
    fileLinks: [],
    uploadFilesProgress: [],
  },
  created() {
    this.$http.get(server_url + "/api/bucket/list").then((response) => {
      this.bucketList = response.data;
      console.log(this.bucketList);
      this.bucketName = this.bucketList[0];
      this.GetFileList();
    });
  },

  methods: {
    GetFileList() {
      this.$http
        .get(server_url + "/api/file/list", {
          params: { bucketName: this.bucketName },
        })
        .then(
          (response) => {
            this.files = response.data;
          },
          (response) => {
            this.files = response.data;
            console.log(response);
          }
        );
    },
    HandleFileChanged() {
      this.$http
        .post(server_url + "/api/file/upload", {
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
    CreateBucket: function () {
      this.$http
        .post(server_url + "/api/bucket/create", {
          bucketName: this.bucket.bucketName,
          bucketRegion: this.bucket.bucketRegion,
          objectLocking: this.bucket.objectLocking,
        })
        .then(
          (response) => {
            this.bucketList = response.data;
          },
          (response) => {
            console.log("err:", response);
          }
        );
      this.createDialog = false;
      // this.bucket = null
    },
    async HandleFileUpload() {
      await this.fileLinks.forEach((fileLink) => {
        let tempProcess = 0;
        let formdata = new FormData();
        this.uploadFiles.find((upFile) => {
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
                    // console.log(uploadFile['fileName'], ":", uploadFile['uploaded']);
                }
              })
            },
          })
          .then(() => {
            console.log("File", fileLink["fileName"], "is uploaded");
            this.uploadFilesProgress.splice(this.uploadFilesProgress.indexOf(fileLink["fileName"]),1)
            this.GetFileList()
          });
      });
      this.uploadDialog = false;
      this.uploadFiles = undefined;
      console.log("test", this.uploadFilesProgress)
    },
    RemoveBucket: function () {
      console.log(this.bucketName);
      this.$http
        .post(server_url + "/api/bucket/remove", {
          bucketName: this.bucketName,
        })
        .then(
          (response) => {
            this.bucketList = response.data;
          },
          (response) => {
            console.log("err:", response);
          }
        );
      this.bucketName = null;
    },
    RemoveFile: async function (filename) {
      await this.$http
        .get(server_url + "/api/file/remove", {
          params: {
            fileName: filename,
            bucketName: this.bucketName,
          },
        })
        .then((response) => {
          this.files = response.data;
        });
      this.GetFileList()
    },
  },
  vuetify: new Vuetify(),
});
