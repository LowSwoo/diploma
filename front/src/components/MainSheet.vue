<template>
  <v-col>
    <v-sheet min-height="70vh" rounded="lg" elevation="4">
      <br />
      <template v-if="files">
        <template v-if="files.length >= 1">
        <v-card class="mx-auto px-1" flat tile>
        <v-list>
          <v-list-item-group v-model="selectedItem" color="primary">
            <v-list-item v-for="(item, i) in files" :key="i">
              <v-list-item-content v-on:click="currentFile = item.name; GetFileInfo()">
                <v-row>
                  <v-col cols="11">
                    <v-list-item-title v-text="item.name"></v-list-item-title>
                  </v-col>
                  <v-col class="text-right">
                    <v-btn icon color="light-blue darken-2" v-bind:href="item.userTags.link">
                      <v-icon>mdi-download</v-icon>
                    </v-btn>
                    <v-btn icon color="red darken-4" @click="RemoveFile(item.name)">
                      <v-icon>mdi-delete-forever</v-icon>
                    </v-btn>
                  </v-col>
                </v-row>
              </v-list-item-content>
            </v-list-item>
          </v-list-item-group>
        </v-list>
      </v-card>
      </template>
    </template>
    <h2 class="text-center" v-else>Для данного расчёта нет данных</h2>
      

      
    </v-sheet>
  </v-col>
</template>
<script>
import settings from "../settings";
export default {
  name: "MainSheet",
  data: () => ({
    bucketName: "",
    currentFile: "",
    files: [],
    selectedItem: null,
  }),
  methods: {
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
        .then((response) => {
            this.files = response.data;
            },
          
        );
    },
  },
  mounted() {
    this.$root.$on("ChangeCurrentBucket", (bucketName) => {
      this.bucketName = bucketName;
      console.log(this.bucketName)
      this.GetFileList();
    });
    this.$root.$on("CheckFilesUpdates", () => { this.GetFileList() })
  },
};
</script>
