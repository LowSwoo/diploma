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
      this.bucketName = this.bucketList[0]
      this.GetFileList()
    });
  },

  methods: {
    GetFileList(){
      this.$http.get("http://localhost:8080/api/file/list", {params: {bucketName: this.bucketName}}).then((response) => {
        this.files = response.data
      },(response) => {this.files = response.data; console.log(response)})
    },
    HandleFileChanged() {
      // console.log(this.uploadFiles.map(el => el.name))
      this.$http
        .post("http://localhost:8080/api/file/upload", {
          bucketName: this.bucketName,
          fileName: this.uploadFiles.map(el => el.name),
        })
        .then((response) => {
          response.data.map(el => this.fileLinks.push({fileName: el["fileName"], url: el["url"]}))
          console.log(this.fileLinks.map(x=>x.fileName));
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
      this.createDialog = false
      // this.bucket = null
    },
    async HandleFileUpload() {
      var formdata = new FormData();
      formdata.append("data", this.uploadFiles);
      this.uploadDialog = false
      await this.$http.put(this.fileLink, formdata, {
        uploadProgress: e => this.uploadStatus = Math.round(e.loaded * 100 / e.total)
      }).then(
        (response) => {
          
        },
        (response) => {
          console.log("err:", response);
        }
        );
      this.uploadFiles = undefined
      this.uploadStatus = 0;
      this.GetFileList()
    },
    RemoveBucket: function () {
      console.log(this.bucketName);
      this.$http
        .post("http://localhost:8080/api/bucket/remove", {
          bucketName: this.bucketName,
        })
        .then(
          (response) => {
            this.bucketList = response.data
          },
          (response) => {
            console.log("err:", response);
          }
        );
        this.bucketName = null
    },
    RemoveFile: function(filename) {
      this.$http.get("localhost:8080/api/file/remove", {params: {
        fileName: filename,
        bucketName: this.bucketName,
      }})
      this.GetFileList()
    }
  },
  vuetify: new Vuetify(),
});