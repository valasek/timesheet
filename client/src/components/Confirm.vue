<!-- Copyright © 2018-2020 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <q-dialog
    v-model="dialog"
    :style="{ zIndex: options.zIndex }"
    :max-width="options.width"
    @keydown.esc="cancel"
    @keydown.enter="agree"
  >
    <q-card>
      <q-toolbar dark :class="options.color" dense flat>
        <q-toolbar-title>{{ title }}</q-toolbar-title>
      </q-toolbar>
      <q-card-section v-show="!!message">
        <!-- eslint-disable-next-line vue/no-v-html -->
        <span v-html="message" />
      </q-card-section>
      <q-card-actions class="pt-0">
        <q-space />
        <q-btn :color="colorAgree" flat="flat" @click.native="agree">Yes</q-btn>
        <q-btn :color="colorCancel" flat="flat" @click.native="cancel">Cancel</q-btn>
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script>
export default {
  data: () => ({
    dialog: false,
    resolve: null,
    reject: null,
    message: null,
    defaultAction: null,
    title: null,
    options: {
      color: 'primary',
      width: 290,
      zIndex: 200
    }
  }),

  computed: {
    colorAgree () {
      if (this.defaultAction === 'agree') {
        return 'primary'
      } else {
        return 'grey'
      }
    },
    colorCancel () {
      if (this.defaultAction === 'cancel') {
        return 'primary'
      } else {
        return 'grey'
      }
    }
  },

  methods: {
    open (title, message, defaultAction, options) {
      this.dialog = true
      this.title = title
      this.message = message
      this.defaultAction = defaultAction
      this.options = Object.assign(this.options, options)
      return new Promise((resolve, reject) => {
        this.resolve = resolve
        this.reject = reject
      })
    },
    agree () {
      this.resolve(true)
      this.dialog = false
    },
    cancel () {
      this.resolve(false)
      this.dialog = false
    }
  }
}
</script>
