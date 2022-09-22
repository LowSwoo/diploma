<template>
  <v-col>
    <v-sheet min-height="70vh" rounded="lg" elevation="4">
      <br />
      <v-card v-for="file in files" tile class="mx-auto">
        <v-list>
          <v-list-item-group color="primary">

            <v-list-item>
              <v-list-item-content>
                <v-row>
                <v-col cols="11">
                  <v-list-item-title>{{file.name}}</v-list-item-title>
                </v-col>
                <v-col class="text-right">
                  <v-icon class="mx-4">mdi-download</v-icon>
                  <v-icon color="red">mdi-close-circle</v-icon>
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
    files: [],
    selectedFile: '',
  }),
  methods: {
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
