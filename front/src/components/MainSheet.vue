<template>
  <v-col>
    <v-sheet min-height="70vh" rounded="lg" elevation="4">
      <br />
      <v-card class="mx-auto px-1" flat  tile>
        <v-list >
          <v-list-item-group v-model="selectedItem" color="primary">
            <v-list-item v-for="(item, i) in files" :key="i">
              <v-list-item-content v-on:click="currentFile = item.name; GetFileInfo()">
                <v-row>
                  <v-col cols="11">
                    <v-list-item-title v-text="item.name"></v-list-item-title>
                  </v-col>
                  <v-col class="text-right">
                    <v-btn icon color="light-blue darken-2"><v-icon>mdi-download</v-icon></v-btn>
                    <v-btn icon color="red darken-4" @click="RemoveFile(item.name)"><v-icon>mdi-delete-forever</v-icon></v-btn>
                  </v-col>
                </v-row>
              </v-list-item-content>
            </v-list-item>
          </v-list-item-group>
        </v-list>
      </v-card>
    </v-sheet>
  </v-col>
</template>
<script>
export default {
  name: "MainSheet",
  data: () => ({
    bucketName: "",
    currentFile: "",
    files: [],
    selectedItem: null,
  }),
  methods: {
    RemoveFile: async function(filename) {
      await this.$http
        .get("http://localhost:8080" + "/api/file/remove", {
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
    GetFileInfo() {
      console.log(this.files.find((x)=>{return x.name == this.currentFile}))
    },
    GetFileList() {
      this.$http
        .get("http://localhost:8080" + "/api/file/list", {
          params: { bucketName: this.bucketName },
        })
        .then(
          (response) => {
            this.files = response.data;
            console.log(this.files);
          },
          (response) => {
            this.files = response.data;
            console.log(response);
          }
        );
    },
  },
  mounted() {
    this.$root.$on("ChangeCurrentBucket", (bucketName) => {
      this.bucketName = bucketName;
      this.GetFileList();
    });
  },
};
</script>
