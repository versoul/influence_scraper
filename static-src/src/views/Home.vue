<template>
  <div class="home">
    <b-form inline class="justify-content-center">
      <b-form-input id="inline-form-input-name"
                    class="mr-sm-3 mb-sm-0"
                    placeholder="Enter Nickname"
                    v-model="nickname">
      </b-form-input>

      <b-button variant="primary"
                :disabled="loading || nickname ===''"
                @click="checkRate">
        <b-icon v-if="loading"
                icon="arrow-clockwise"
                animation="spin">
        </b-icon>
        <template v-else>
          Check Rate
        </template>
      </b-button>
    </b-form>
    <div class="rate text-center"
         v-if="rate !== ''">
      <b>{{nickname}}</b> engagement rate is: <b>{{rate}}%</b><br/>
      <i>Your formula (Engagement Rate = (likes + comments) / followers) <b>is wrong</b> look FIXME comments</i>
    </div>
  </div>
</template>

<script>
import { CheckRate } from '@/util/api'

export default {
  name: 'Home',
  data () {
    return {
      nickname: 'thetwobohemians',
      rate: '',
      loading: false
    }
  },
  methods: {
    async checkRate () {
      this.loading = true
      this.rate = ''
      try {
        const resp = await CheckRate(this.nickname)
        this.rate = resp.data
      } catch (e) {
        console.log('ERR', e)
      }
      this.loading = false
    }
  }
}
</script>

<style scoped lang="scss">
.home{
  form {
    button {
      width: 110px;
    }
  }
  .rate {
    margin-top: 20px;
  }
}
</style>
