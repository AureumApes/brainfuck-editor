<script lang="ts">
    function onSave(){
        let code = document.getElementById("code-input").value
        let name = document.getElementById("name").value
        fetch("/save?name=" + name, {
            method: "post",
            body: code
        }).then()
    }
</script>

<button on:click|preventDefault={onSave}>Save</button>