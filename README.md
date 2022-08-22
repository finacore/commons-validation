# Commons Validation

This module encapsulate some methods and data structure responsibles to perform and manage the erors that may be was generated by the validation execution.

The main objetive to create this package is to assure that the validation result every follow the same pattern, in all seervice of thei project.

As validation language, the __en_US__ was choosen and all message will be returned in the refered language.

## Usage

The fascicle below shows the use of this library.

```go
import commonsvalidation "github.com/finacore/commons-validation"

errs := commonsvalidation.ValidateModel(theModelDataStructure)
```

As the response, the method will give you a nil return, that means that the model has no error, ou a __commonserrors.ValidationError__ array containing all erros detected.