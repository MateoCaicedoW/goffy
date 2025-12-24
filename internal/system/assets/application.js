import { Application } from "stimulus"
import ConverterController from "converter"
window.Stimulus = Application.start()
Stimulus.register("converter", ConverterController)