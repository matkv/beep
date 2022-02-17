use tray_item::TrayItem;

fn main() {
    create_tray_icon();
}

fn create_tray_icon() {
    gtk::init().unwrap();
}
