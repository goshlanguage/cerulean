<p align="center">
    <a href="https://aka.ms/free-account">
    <img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/296b4d47f5313822b348e442837ca2d32a7704a3/Azure_Gophers.png" width="360"></a>
</p>

# cerulean
Cerulean is a golang library for mocking the [azure-sdk-for-go](https://github.com/Azure/azure-sdk-for-go).


# Design
Cerulean aims to mimic the codepaths of the azure-sdk-for-go for ease of use, and clarity when creating mocks. The developer should be able to copy the imports from their code, and replace `azure/azure-sdk-for-go` with `goshlanguage/cerulean` in order to clearly identify how to import the right mock(s).
