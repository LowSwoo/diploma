<template>
  <v-col cols="2">
    <v-sheet rounded="lg" elevation="2">
      <v-list color="transparent">
        <v-list-item>
            <v-list-item>Текущие расчёты</v-list-item>
        </v-list-item>
        <v-divider class="my-2"></v-divider>

        <v-list-item v-for="bucket in bucketList" :key="bucket" link>
          <v-list-item-content v-on:click="$root.$emit('ChangeCurrentBucket', bucket)">
            <v-list-item-title> {{ bucket }} </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-sheet>
  </v-col>
</template>

<script>

export default {
    name: "LeftNav",
    data: () =>({
      bucketList: [],
      bucketName: '',
    }),
    methods: {
      GetBucketList() {
        this.$http.get("http://localhost:8080/api/bucket/list").then((resp) => {
        this.bucketList = resp.data;
        this.bucketName = this.bucketList[0]
        this.$root.$emit('ChangeCurrentBucket', this.bucketName)
      })
      }
    },
    created() {
      this.GetBucketList()
    },
    mounted() {
      this.$root.$on('UpdateBucketList', () => this.GetBucketList())
    },

}


</script>
