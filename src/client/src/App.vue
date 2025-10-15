<script setup lang="ts">
import { createClient } from '@connectrpc/connect'
import { GreetingService } from '../generated/helloapp/v1/hello_pb'
import { createConnectTransport } from '@connectrpc/connect-web'
import { ref } from 'vue'

const transport = createConnectTransport({
  baseUrl: 'http://localhost:8080',
})
const client = createClient(GreetingService, transport)

const response = ref<string>('')
const errContent = ref<string>('')
async function sayHello() {
  try {
    const res = await client.hello({ name: 'Ken' })
    response.value = res.message
  } catch (err) {
    errContent.value = 'Helloメソッドの実行に失敗しました。'
    console.log(err)
  }
}
</script>

<template>
  <button @click="sayHello">gRPCメソッド呼び出し</button>
  <p>{{ response }}</p>
  <p>{{ errContent }}</p>
</template>

<style scoped></style>
