new Vue({
  el: "#app",
  data: {
    bucket: {
      bucketName: "",
      bucketRegion: "",
      objectLocking: false,
    },
    uploadStatus: 40,
    bucketName: undefined,
    bucketList: [],
    createDialog: false,
    uploadDialog: false,
    uploadSnackbar: false,
    uploadFiles: [],
    files: [],
    file: null,
    fileLinks: [],
  },
  created() {
    this.$http.get("http://localhost:8080/api/bucket/list").then((response) => {
      this.bucketList = response.data;
      console.log(this.bucketList);
      this.bucketName = this.bucketList[0];
      this.GetFileList();
    });
  },

  methods: {
    GetFileList() {
      this.$http
        .get("http://localhost:8080/api/file/list", {
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
        .post("http://localhost:8080/api/file/upload", {
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
        .post("http://localhost:8080/api/bucket/create", {
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
      this.fileLinks.forEach((x) => {
        let arr = [];
        let formdata = new FormData();

        this.uploadFiles.find((y) => {
          if (x["fileName"] == y["name"]) {
            formdata.append("data", y);
          }
        });

        this.$http.put(x["url"], formdata)
      });
      // formdata.append("data", this.uploadFiles);
      // console.log(formdata)
      this.uploadDialog = false;
      // this.$http.put(this.fileLink, formdata, {
      //   uploadProgress: e => this.uploadStatus = Math.round(e.loaded * 100 / e.total)
      // })
      this.uploadFiles = undefined;
      this.uploadStatus = 0;
      this.GetFileList();
    },
    RemoveBucket: function () {
      console.log(this.bucketName);
      this.$http
        .post("http://localhost:8080/api/bucket/remove", {
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
    RemoveFile: function (filename) {
      this.$http
        .get("http://localhost:8080/api/file/remove", {
          params: {
            fileName: filename,
            bucketName: this.bucketName,
          },
        })
        .then((response) => {
          
          this.files = response.data
        });
      // this.GetFileList()
    },
  },
  vuetify: new Vuetify(),
});
