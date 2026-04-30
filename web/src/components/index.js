import LyPictureMultipleUpload from '@/components/upload/mutiple-pictures.vue'
import LyPictureSingleUpload from '@/components/upload/single-picture.vue'
import LyFileMultipleUpload from '@/components/upload/lyfileUpload.vue'
import LyTeditor from '@/components/teditorjs/lyTEditor.vue'

export default {
    install(app) {
        app.component("ly-public-pictrue-multiple-upload", LyPictureMultipleUpload);
        app.component("ly-public-pictrue-single-upload", LyPictureSingleUpload);
        app.component("ly-public-file-multiple-upload", LyFileMultipleUpload);
        app.component("ly-public-teditor", LyTeditor);
    },
};