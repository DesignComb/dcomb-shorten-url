<script setup lang="ts">
import 'vue-cropper/dist/index.css'
import {VueCropper} from "vue-cropper"
import {ref, reactive} from 'vue'

let uploadedUrl = ref('')
let selectedFile = reactive({
    fileObj: ''
})
let loading = ref(false)
let preview = reactive({image: ''})

const option = {
    img: '', // 裁剪圖片的地址
    size: 1, // 裁剪生成圖片的品質
    full: true, // 是否輸出原圖比例的截圖 默認false
    outputType: 'png', // 裁剪生成圖片的格式 默認jpg
    canMove: false, // 上傳圖片是否可以移動
    fixedBox: false, // 固定截圖框大小 不允許改變
    original: false, // 上傳圖片按照原始比例渲染
    canMoveBox: true, // 截圖框能否拖動
    autoCrop: true, // 是否默認生成截圖框
    // 只有自動截圖開啟 寬度高度才生效
    autoCropWidth: 1230, // 默認生成截圖框寬度
    autoCropHeight: 360, // 默認生成截圖框高度
    centerBox: true, // 截圖框是否被限制在圖片裡面
    high: false, // 是否按照設備的dpr 輸出等比例圖片
    enlarge: 1, // 圖片根據截圖框輸出比例倍數
    mode: '100% auto', // 圖片默認渲染方式
    maxImgSize: 4000, // 限制圖片最大寬度和高度
    limitMinSize: [615, 180], // 更新裁剪框最小屬性
    infoTrue: false, // true 為展示真實輸出圖片寬高 false 展示看到的截圖框寬高
    fixed: true, // 是否開啟截圖框寬高固定比例  (默認:true)
    fixedNumber: [1, 1] // 截圖框的寬高比例
}

const realTime = (data: any) => {
    preview.image = data
}

const onFileSelectEvent = (e: any) => {
    if (!e.target.files.length) {
        loading.value = false;
        return;
    }

    const file = e.target.files[0];
    if (typeof FileReader === "function") {
        const res = readAsDataURL(file);
        res.then((res: any) => {
            // const originImage = new Image();
            // originImage.src = res;
            // this.selectedFile = originImage.outerHTML;

            selectedFile.fileObj = res
            console.log(selectedFile)
            console.log(res)
            // this.$refs.cropper.replace(this.selectedFile);
        });
    } else {
        console.error("Sorry, FileReader API not supported");
    }
}

const readAsDataURL = (file: any) => {
    return new Promise((resolve, reject) => {
        const fr = new FileReader();
        fr.onerror = reject;
        fr.onload = function (event) {
            resolve(fr.result);
        };
        fr.readAsDataURL(file);
    });
}

</script>
<template>
    <div class="relative">
        <!--    <img v-if="selectedFile" :src="selectedFile" width="200"  alt=""/>-->
        <input id="imageInput" ref="FileInput" class="hidden" type="file" @input="onFileSelectEvent"/>
        <label v-if="!selectedFile.fileObj" for="imageInput"
               class="absolute absolute-middle w-fit h-fit text-4xl z-50 text-teal-500
                border rounded-full bg-gray-800 shadow-md p-4 hover:opacity-80 cursor-pointer">
            <i class="bx bx-upload"></i>
        </label>
        <div class="cropper-container">
            <div class="cropper-el">
                <vueCropper
                        ref="cropper"
                        :img="selectedFile.fileObj"
                        :output-size="option.size"
                        :output-type="option.outputType"
                        :info="true"
                        :full="option.full"
                        :can-move="option.canMove"
                        :can-move-box="option.canMoveBox"
                        :fixed-box="option.fixedBox"
                        :original="option.original"
                        :auto-crop="option.autoCrop"
                        :auto-crop-width="option.autoCropWidth"
                        :auto-crop-height="option.autoCropHeight"
                        :center-box="option.centerBox"
                        :high="option.high"
                        :info-true="option.infoTrue"
                        @realTime="realTime"
                        :enlarge="option.enlarge"
                        :fixed="option.fixed"
                        :fixed-number="option.fixedNumber"
                />
            </div>
        </div>
        <!--    {{ preview }}-->
        <!--    <div :style="{width: preview?.image?.div?.width, height: preview?.image?.div?.height}"-->
        <!--         style="overflow: hidden;border: solid 2px red">-->
        <!--      <img-->
        <!--          alt=""-->
        <!--          :src="preview?.image?.url"-->
        <!--          :style="preview?.image?.img"-->
        <!--      />-->
        <!--    </div>-->
    </div>
</template>

<style>
.cropper-el {
    height: 250px;
    width: 100%;
}

.cropper-container {
    display: flex;
    justify-content: center;
}

</style>