<template>
  <v-dialog v-model="showCreateDialog" persistent max-width="600px">
    <v-card>
      <v-card-title>
        <span class="text-h5">Создать расчёт</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12">
              <v-text-field
                label="Название расчёта"
                v-model="bucket.name"
              ></v-text-field>
            </v-col>
            <v-col cols="12">
              <v-text-field label="Bucket region (пока бд не прикрутил - путь так)" v-model="bucket.region">
              </v-text-field>
            </v-col>
            <v-switch
              :label="`Object locking: ${bucket.objectLocking.toString()}`"
              v-model="bucket.objectLocking"
            >
            </v-switch>
          </v-row>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
          color="blue darken-1"
          text
          v-on:click="$emit('CloseCreateDialog')"
        >
          Close
        </v-btn>
        <v-btn color="blue darken-1" text @click="CreateBucket"> Create </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import settings from '@/settings';
export default {
  name: "CreateBucket",
  props: ["showCreateDialog"],
  data: () => ({
    bucket: {
      name: "",
      region: "",
      objectLocking: false,
    },
  }),
  methods: {
    CreateBucket() {
      this.$http
        .post(settings.ebaHOST + "api/bucket/create", {
          bucketName: this.bucket.name,
          bucketRegion: this.bucket.region,
          objectLocking: this.bucket.objectLocking,
        })
        .then(
          () => {
             this.$root.$emit('UpdateBucketList')
          },
          (response) => {
            console.log("err:", response);
          }
        );
        this.bucket = {
          name: '',
          region: '',
          objectLocking: false,
        }
        this.$emit('CloseCreateDialog')
    },
  },
};
</script>
