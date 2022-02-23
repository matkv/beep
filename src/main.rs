use std::{thread, time::Duration};

use rodio::{
    source::{Amplify, SineWave, TakeDuration},
    OutputStream, Sink, Source,
};
use tray_item::TrayItem;

fn main() {
    setup_tray_item();
}

fn setup_tray_item() {
    gtk::init().unwrap();

    let mut tray = TrayItem::new("Beep", "accessories-calculator").unwrap();

    tray.add_menu_item("Beep once", || {
        play_sound();
    })
    .unwrap();

    tray.add_menu_item("Beep twice", || {
        for _number in 0..2 {
            play_sound();
            thread::sleep(Duration::from_millis(500));
        }
    })
    .unwrap();

    tray.add_menu_item("Quit", || {
        gtk::main_quit();
    })
    .unwrap();

    gtk::main();
}

fn play_sound() {
    let sinewave = setup_sinewave();

    let (_stream, stream_handle) = OutputStream::try_default().unwrap();
    let sink = Sink::try_new(&stream_handle).unwrap();
    sink.append(sinewave);
    sink.sleep_until_end();
}

fn setup_sinewave() -> Amplify<TakeDuration<SineWave>> {
    let sine = SineWave::new(340.0)
        .take_duration(Duration::from_secs_f32(0.75))
        .amplify(0.75);

    return sine;
}
