<script lang="ts">

    function onBuild(){
        let code = document.getElementById("code-input").value
        fetch("/build", {
            method: "post",
            body: code
        }).then()
        console.log(code)
    }
</script>

<button id="build" on:click|preventDefault={onBuild}>
    Build
</button>