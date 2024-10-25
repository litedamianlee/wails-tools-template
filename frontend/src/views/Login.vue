<script lang="ts" setup>
// import { login } from '../apis/modules/public';
import {reactive, ref} from "vue";  // Adjust the path as needed

// import {login} from "../apis/login";
// import {Auth} from "../store/store.ts";
import router from "../router.ts";
import {Auth} from "../store/store.ts";
// import {ElMessage} from "element-plus";

// 输入登录参数
const loginParams = reactive({account: '', password: ''});
// 错误提示
const errMsg = ref('');

// 登录
async function submitLogin() {
  if (loginParams.account.length < 5 || loginParams.account.length > 12) {
    errMsg.value = '用户名长度应在 5 至 12 位之间';
    return;
  }
  if (loginParams.password.length < 6 || loginParams.password.length > 18) {
    errMsg.value = '密码长度应在 6 至 18 位之间';
    return;
  }
  // const res = await login({
  //   account: loginParams.account,
  //   password: loginParams.password,
  // })
  Auth.logged = true; // 设置登录状态
  router.push('/home');
  // if (res.code === 200) {
  //   let token = res.data['access_token']
  //   localStorage.setItem("token", token); // 将 token 存储到本地
  //   ElMessage.success("登陆成功")
  //   setTimeout(function () {
  //     Auth.logged = true; // 设置登录状态
  //     router.push('/home');
  //   }, 1500);
  // } else {
  //   ElMessage.warning( res.message)
  // }
}
</script>

<template>
  <section style="--wails-draggable:drag" class="w-full h-screen">
    <div class="pt-36">
      <div class="mx-auto w-96 space-y-6 bg-white p-12 rounded-lg">
        <h1 class="text-center text-2xl font-semibold dark:text-black">登录</h1>
        <input v-model="loginParams.account" type="text" placeholder="用户名" class="inputClass"
               @input="()=> errMsg = ''"/>
        <input v-model="loginParams.password" type="password" placeholder="密码" class="inputClass"
               @input="()=> errMsg = ''"/>
        <button class="btnClass" @click="submitLogin">登录</button>
        <p v-show="errMsg" class="text-sm text-center text-red-500">{{ errMsg }}</p>
      </div>
    </div>
  </section>
</template>

<style scoped lang="postcss">
.inputClass {
  @apply py-3 px-4  w-full border border-gray-200 rounded-md text-sm focus:border-blue-500 focus:ring-blue-500 dark:bg-slate-200/50 dark:border-gray-400 dark:text-gray-700;
}

.btnClass {
  @apply w-full py-3 px-4 bg-blue-500 text-white rounded-md text-sm hover:bg-blue-600;
}
</style>
