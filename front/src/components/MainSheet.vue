<template>
  <v-col>
    <v-card min-height="70vh" rounded="lg" elevation="4">
      <!-- <br /> -->
      <!-- <v-toolbar> -->
      <template v-if="folders">
        <v-tabs grow>
          <v-tab v-for="folder in folders" :key="folder" @click="changeFolder(folder)">{{ folder }}</v-tab>
        </v-tabs>
        <!-- </v-toolbar> -->
      </template>
      <template v-if="files">
        <v-card class="mx-auto px-1" flat tile>
          <v-list>
            <v-list-item v-for="item in Object.keys(data[currentFolder]['files'])" :key="item">
              <v-list-item-content>
                <v-row>
                  <v-col cols="11">
                  <v-list-item-title v-text="item"></v-list-item-title>
                  <!-- {{files[item]['url']}} -->
                  </v-col>
                  <v-col class="text-right">
                    <v-btn icon color="blue darken-2" v-bind:href="files[item]['url']">
                      <v-icon>mdi-download</v-icon>
                    </v-btn>
                    <v-btn icon color="red darken-4" @click="RemoveFile(item)"><v-icon>mdi-delete-forever</v-icon></v-btn>
                  </v-col>
                </v-row>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-card>
      </template>
      <h2 class="text-center" v-else>Для данного расчёта нет данных</h2>



    </v-card>
  </v-col>
</template>
<script>
import settings from "../settings";
export default {
  name: "MainSheet",
  data: () => ({
    bucketName: "",
    currentFile: "",
    files: null,
    selectedItem: null,
    currentFolder: "",
    folders: [],
    data: null,
  }),
  model: {
    prop: 'currentFolder',
    event: 'click',
  },
  methods: {
    changeFolder(folder) {
      this.currentFolder = folder
      this.files = this.data[this.currentFolder]["files"]
      this.$emit('click', this.currentFolder)
    },
    RemoveFile: async function (filename) {
      await this.$http
        .get(settings.ebaHOST + "api/file/remove", {
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
      console.log(this.files.find((x) => { return x.name == this.currentFile }))
    },
    GetFileList() {
      this.$http
        .get(settings.ebaHOST + "api/file/list", {
          params: { bucketName: this.bucketName },
        })
        .then(response => {
          this.data = response.data
          this.folders = Object.keys(this.data)
          this.currentFolder = this.folders[0]
          this.files = this.data[this.currentFolder]["files"]
        
        })

    },
  },
  mounted() {
    this.GetFileList()
    // this.changeFolder(this.folders[0])
    this.$emit('click', this.folders[0])
    this.$root.$on("ChangeCurrentBucket", (bucketName) => {
      this.bucketName = bucketName;
      console.log(this.bucketName)
      this.GetFileList();
    });
    this.$root.$on("CheckFilesUpdates", () => { this.GetFileList() })
  },
};
</script>
