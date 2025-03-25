
<script lang="ts">
    import Pocketbase from "pocketbase"
    import { onDestroy, onMount } from "svelte";

    import { AnsiUp } from "ansi_up"
    const ansiup = new AnsiUp()

    const pb = new Pocketbase("")
    // const uri = document.baseURI

    let log = $state("")

    onMount(async () => {
        const res = await fetch("http://localhost:8090/api/v1/docker/c9f3a113e3df/log")
        const reader = res.body!.getReader()
        const decoder = new TextDecoder()
        while (true) {
            const {done, value} = await reader.read()
            if (done) {
                break
            }
            const text = decoder.decode(value)
            log = log + ansiup.ansi_to_html(text)
        }
    })

    onDestroy(() => {
        pb.realtime.unsubscribe("docker.container.log.c9f3a113e3df")
    })

</script>

<main>
    {@html "<p style='white-space: pre-line'>" + log + "<p/>"}
</main>