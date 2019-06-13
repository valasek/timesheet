<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <q-dialog v-model="dialog" :max-width="options.width" :style="{ zIndex: options.zIndex }" @keydown.esc="cancel" @keydown.enter="agree">
    <q-card>
      <q-toolbar dark :class="options.color" dense flat>
        <q-toolbar-title>
          {{ title }}
        </q-toolbar-title>
      </q-toolbar>
      <q-card-section v-show="!!message">
        {{ message }}
      </q-card-section>
      <q-card-actions class="pt-0">
        <q-space />
        <q-btn color="primary" flat="flat" @click.native="agree">
          Yes
        </q-btn>
        <q-btn color="grey" flat="flat" @click.native="cancel">
          Cancel
        </q-btn>
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
    title: null,
    options: {
      color: 'primary',
      width: 290,
      zIndex: 200
    }
  }),
  methods: {
    open (title, message, options) {
      this.dialog = true
      this.title = title
      this.message = message
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
