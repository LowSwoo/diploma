<template>
    <v-app-bar app elevation="2" color="topNav">
        <v-app-bar-nav-icon @click.stop='toggleNavigationDrawer'></v-app-bar-nav-icon>
        <v-container class="py-0 fill-height">
            <!-- <v-btn fab small>
              <v-avatar class="mr-10" color="grey darken-3" size="32" >L</v-avatar>
            </v-btn> -->

            <v-btn text @click="showCreateDialog = true"> Создать расчёт</v-btn>
            <v-btn text @click="showCreateDescriptionDialog = true">Посмотреть описание</v-btn>


            <v-spacer></v-spacer>

            <v-btn width="400" text @click="uploadSnackbar = true" v-if="uploadFilesProgress.length !== 0">
                <v-progress-linear
                        v-bind:value="uploadFilesProgress.reduce((sum,elem)=>{return sum + elem.uploaded},0)/uploadFilesProgress.length"></v-progress-linear>
            </v-btn>
            <v-snackbar v-model="uploadSnackbar" right top app>
                <v-row v-for="el in uploadFilesProgress">
                    <v-col>{{ el.fileName }}</v-col>
                    <v-col>{{ el.bucketName }}</v-col>
                    <v-col
                    >
                        <v-progress-circular :value="el.uploaded"></v-progress-circular
                        >
                    </v-col>
                </v-row>
                <template v-slot:action="{ attrs }">
                    <v-btn
                            color="blue"
                            text
                            v-bind="attrs"
                            @click="uploadSnackbar = false"
                            bottom
                    >Close
                    </v-btn
                    >
                </template>
            </v-snackbar>
            <v-btn v-if="bucketName" icon @click="showUploadDialog = true">
                <v-icon>mdi-cloud-upload</v-icon>
            </v-btn>


            <v-btn v-if="bucketName" class="outlined x-small elevation-0 " @click="RemoveBucket" text>
                {{ bucketName }}
                <v-icon color="red darken-4" class="mx-2">mdi-delete-forever</v-icon>
            </v-btn>
            <CreateBucket :showCreateDialog="showCreateDialog"
                          v-on:CloseCreateDialog="showCreateDialog = false"></CreateBucket>
            <UploadFile :currentFolder="currentFolder" :bucketName="bucketName" :showUploadDialog="showUploadDialog"
                        v-on:CloseUploadDialog="showUploadDialog = false"
                        v-on:uploadProgress="uploadProgress"></UploadFile>
        <DescriptionView :showDescription='showCreateDescriptionDialog'/>

        </v-container>
    </v-app-bar>
</template>
<script>
import CreateBucket from "./dialogs/CreateBucket.vue";
import UploadFile from "./dialogs/UploadFile.vue";
import Settings from "../settings"
import DescriptionView from "@/components/dialogs/DescriptionView.vue";

export default {
    data: () => ({
        uploadSnackbar: false,
        showCreateDialog: false,
        showUploadDialog: false,
        showCreateDescriptionDialog: false,
        bucketName: '',
        uploadFilesProgress: [],
    }),
    props: ['currentFolder'],
    methods: {
        uploadProgress(data) {
            this.uploadFilesProgress = data
        },
        toggleNavigationDrawer() {
            this.$emit('toggleNavigationDrawer')
        },
        RemoveBucket() {
            this.$http
                .post(Settings.ebaHOST + "api/bucket/remove", {
                    bucketName: this.bucketName,
                })
                .then(
                    (response) => {
                        this.$root.$emit("UpdateBucketList")
                    },
                    (response) => {
                        console.log("err:", response);
                    }
                );
        }
    },
    mounted() {
        this.$root.$on('ChangeCurrentBucket', (bucketName) => this.bucketName = bucketName)
    },
    components: {CreateBucket, UploadFile, DescriptionView},
};
</script>
