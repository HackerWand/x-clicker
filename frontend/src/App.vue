<style scoped lang="less"></style>

<template>
  <n-card size="small">
    <n-form :disabled="data.Running">
      <n-form-item label="点击类型">
        <n-radio-group @update:value="setConfig" v-model:value="data.Type">
          <n-radio-button value="left">左键</n-radio-button>
          <n-radio-button value="right">右键</n-radio-button>
        </n-radio-group>
      </n-form-item>
      <n-form-item label="点击方式">
        <n-radio-group @update:value="setConfig" v-model:value="data.Double">
          <n-radio-button :value="false">单击</n-radio-button>
          <n-radio-button :value="true">双击</n-radio-button>
        </n-radio-group>
      </n-form-item>
      <n-form-item label="速度">
        <n-input-number @update:value="setConfig" :precision="0" v-model:value="data.Speed" :min="100">
          <template #suffix>毫秒</template>
        </n-input-number>
      </n-form-item>
      <n-form-item label="启动热键">
        <n-select @update:value="setConfig" :options="hotKeys.filter(v => v.value !== data.StopKey)" v-model:value="data.StartKey"></n-select>
      </n-form-item>
      <n-form-item label="关闭热键">
        <n-select @update:value="setConfig" :options="hotKeys.filter(v => v.value !== data.StartKey)" v-model:value="data.StopKey"></n-select>
      </n-form-item>
      <n-form-item :show-label="false">
        <n-space>
          <n-button @click="start" :loading="data.Running" type="info">开始</n-button>
          <n-button @click="stop" :disabled="!data.Running" type="warning">停止</n-button>
        </n-space>
      </n-form-item>
    </n-form>
  </n-card>
</template>

<script lang="ts" setup>
import { NCard, NForm, NFormItem, NSpace, NRadioGroup, NRadioButton, NInputNumber, NButton, NSelect } from 'naive-ui'
import { ref } from 'vue'
import { Start, Stop, SetConfig, GetConfig } from '../wailsjs/go/main/App'

type Config = {
  Type: 'left' | 'right',
  Double: boolean,
  Speed: number,
  StartKey: string,
  StopKey: string,
  Running: boolean
}

const data = ref<Partial<Config>>({})

const hotKeys = [
  { label: 'F1', value: 'F1' },
  { label: 'F2', value: 'F2' },
  { label: 'F3', value: 'F3' },
  { label: 'F4', value: 'F4' },
  { label: 'F5', value: 'F5' },
  { label: 'F6', value: 'F6' },
  { label: 'F7', value: 'F7' },
  { label: 'F8', value: 'F8' },
  { label: 'F9', value: 'F9' },
  { label: 'F10', value: 'F10' },
  { label: 'F11', value: 'F11' },
  { label: 'F12', value: 'F12' }
]

function start () {
  data.value.Running = true
  Start()
}

function stop () {
  data.value.Running = false
  Stop()
}

let timer: number
function setConfig () {
  clearTimeout(timer)
  timer = window.setTimeout(() => {
    SetConfig(data.value)
  }, 500)
}

GetConfig().then(res => {
  data.value = res
})
</script>
