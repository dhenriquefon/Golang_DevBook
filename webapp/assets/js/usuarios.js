
$('#parar-de-seguir').on('click', pararDeSeguir)
$('#seguir').on('click', seguir)
$('#editar-usuario').on('submit', editarUsuario)
$('#atualizar-senha').on('submit', atualizarSenha)
$('#deletar-usuario').on('click', deletarUsuario)

function pararDeSeguir() {
    const usuarioID = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioID}/parar-de-seguir`,
        method: "POST",        
    }).done(function() {
        window.location = `/usuarios/${usuarioID}`
    }).fail(function(){
        Swal.fire("Ops...", "Erro ao parar de seguir usuario", "error")
        $('#parar-de-seguir').prop('disabled', false)
    })

}

function seguir() {
    const usuarioID = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioID}/seguir`,
        method: "POST",        
    }).done(function() {
        window.location = `/usuarios/${usuarioID}`
    }).fail(function(){
        Swal.fire("Ops...", "Erro ao seguir usuario", "error")
        $('#seguir').prop('disabled', false)
    })

}

function editarUsuario(evento) {
    evento.preventDefault()

    $.ajax({
        url:"/editar-usuario",
        method: "PUT",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val()
        }
    }).done(function() {
        Swal.fire("Sucesso!", "Usuario atualizado com sucesso", "success")
            .then(function(){
                window.location = "/perfil"
            });
    }).fail(function(){
        Swal.fire("Ops...", "Erro ao atualizar usuario", "error")
    })

}

function atualizarSenha(evento) {
    evento.preventDefault()

    if ($('#nova-senha').val() != $('#confirmar-senha').val()) {
        Swal.fire("Ops...", "As senhas não coincidem!", "error")
        return;
    }

    $.ajax({
        url:"/atualizar-senha",
        method: "POST",
        data: {
            atual: $('#senha-atual').val(),
            nova: $('#nova-senha').val()
        }
    }).done(function() {
        Swal.fire("Sucesso!", "Senha atualizado com sucesso", "success")
            .then(function(){
                window.location = "/perfil"
            });
    }).fail(function(){
        Swal.fire("Ops...", "Erro ao atualizar senha do usuario", "error")
    })

}

function deletarUsuario(evento) {
    evento.preventDefault()

    Swal.fire( {
        title: "Atenção!",
        text: "Tem certeza que deseja excluir seu usuario? Ação irreversível",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao) {
        if (!confirmacao.value) return;

        $.ajax({
            url:"/deletar-usuario",
            method: "DELETE"
        }).done(function() {
            Swal.fire("Sucesso!", "Usuario excluido com sucesso", "success")
                .then(function(){
                    window.location = "/logout"
                });
        }).fail(function(){
            Swal.fire("Ops...", "Erro ao excluir usuario usuario", "error")
        })    
    })

}