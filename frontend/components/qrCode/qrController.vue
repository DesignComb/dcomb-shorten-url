<script setup lang="ts">
import VueQrcode from '@chenfengyuan/vue-qrcode';

const props = defineProps(['completeUrl']);

let qrUrl = ref('')
const onDataUrlChange = (dataUrl: string) => {
  qrUrl.value = dataUrl
}

import html2canvas from 'html2canvas'

const downloadQRCode = () => {
  const qrCodeElement = document.querySelector(`#QrCode`) as HTMLElement
  html2canvas(qrCodeElement, {
    useCORS: true,
    allowTaint: false,
    logging: true
  })
      .then(canvas => {
        let link = document.getElementById('link') as HTMLElement
        link.setAttribute('download', 'dco-QRcode.png')
        link.setAttribute('href', canvas.toDataURL("image/png").replace("image/png", "image/octet-stream"))
        link.click()
      })
}
</script>

<template>
  <div class="flex w-full justify-between items-end">
    <vue-qrcode
        id="QrCode"
        :value="completeUrl"
        :options="{ width: 150 ,margin: 2}"
        @change="onDataUrlChange"
    />
    <a id="link"></a>
    <button class="border border-teal-500 rounded-full px-12 py-2 text-white flex justify-between items-center
                  hover:bg-gray-700"
        @click="downloadQRCode">Download
      <i class="bx bx-download pl-1"></i>
    </button>
  </div>
</template>

<style scoped>
</style>