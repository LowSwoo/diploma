<template>
  <v-snackbar v-model="uploadSnackbar" light right top app>
    <v-row v-for="el in uploadFilesProgress">
      <v-col>{{ el.fileName }}</v-col>
      <v-col>{{ el.bucketName }}</v-col>
      <v-col
        ><v-progress-circular :value="el.uploaded"></v-progress-circular
      ></v-col>
    </v-row>
    <template v-slot:action="{ attrs }">
      <v-btn
        color="blue"
        text
        v-bind="attrs"
        @click="uploadSnackbar = false"
        bottom
        >Close</v-btn
      >
    </template>
  </v-snackbar>
</template>

<script>
    export default {
    name: 'UploadSnackbar'
}

</script>
