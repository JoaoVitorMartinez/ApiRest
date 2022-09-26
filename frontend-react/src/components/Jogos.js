import React, { Component } from 'react';
import axios from 'axios';
import '../components/Jogos.css';

export default class Jogos extends Component {
    state = {
        Jogos: []
    }

    componentDidMount() {
        axios.get('http://localhost:8000/api/jogos')
            .then(res => {
                const Jogos = res.data;
                this.setState({ Jogos })
            })
    }

    render() {
        return (
            <div class="container">
                <div class="table-responsive">
                    <table class="table table-dark">
                        <thead>
                            <tr scope="row">
                                <th scope="col">Id</th>
                                <th scope="col">Nome</th>
                                <th scope="col">Descriçao</th>
                                <th scope="col">Preço</th>
                                <th scope="col">Nota</th>
                                
                                
                            </tr>
                        </thead>
                        <tbody>
                    
                        {this.state.Jogos.map((j, id )=>
                        
                        
                                <tr key={id} >
                                    <th key={id} scope="row">
                                        <p>{j.Id}</p>
                                    </th>
                                    <td >
                                        <h2>{j.Nome}</h2>
                                    </td>
                                    <td>
                                        <p>{j.Descricao}</p>
                                    </td>
                                    <td>
                                        <p>R${j.Preco}</p>
                                    </td>
                                    <td>
                                        <p>{j.Nota}</p>
                                    </td>
                                </tr>
                            )}
                            </tbody>
                        </table>
                    </div>
            </div>
        );
    }
}